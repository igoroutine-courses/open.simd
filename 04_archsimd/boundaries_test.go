package contains

import (
	"fmt"
	"slices"
	"testing"
)

func TestEveryLaneAndTail(t *testing.T) {
	t.Run("int8", func(t *testing.T) { checkContains(t, SliceContainsInt8) })
	t.Run("int32", func(t *testing.T) { checkContains(t, SliceContainsInt32) })
	t.Run("int64", func(t *testing.T) { checkContains(t, SliceContainsInt64) })
}

func checkContains[T ~int8 | ~int32 | ~int64](t *testing.T, contains func([]T, T) bool) {
	t.Helper()
	if contains(nil, 0) {
		t.Fatal("nil input must not contain zero")
	}
	for n := 0; n <= 129; n++ {
		for _, target := range []T{0, -1, 127, -128} {
			// Targets just outside an unaligned, capacity-limited input must
			// not be considered matches. Each in-range lane is tested too.
			backing := make([]T, n+2)
			backing[0], backing[n+1] = target, target
			s := backing[1 : n+1 : n+1]
			for i := range s {
				s[i] = target + 1
			}
			if contains(s, target) {
				t.Fatalf("false match: n=%d target=%d", n, target)
			}
			for i := range s {
				s[i] = target
				before := slices.Clone(backing)
				if got := contains(s, target); got != slices.Contains(s, target) {
					t.Fatalf("missed lane: n=%d index=%d target=%d", n, i, target)
				}
				if !slices.Equal(backing, before) {
					t.Fatalf("input modified: n=%d index=%d", n, i)
				}
				s[i] = target + 1
			}
		}
	}
}

func BenchmarkContains(b *testing.B) {
	b.Run("int8", func(b *testing.B) { benchmarkContains(b, SliceContainsInt8) })
	b.Run("int32", func(b *testing.B) { benchmarkContains(b, SliceContainsInt32) })
	b.Run("int64", func(b *testing.B) { benchmarkContains(b, SliceContainsInt64) })
}

func benchmarkContains[T ~int8 | ~int32 | ~int64](b *testing.B, contains func([]T, T) bool) {
	for _, n := range []int{19, 1024, 100_000} {
		for _, hit := range []string{"missing", "first", "last"} {
			s := make([]T, n)
			if hit == "first" {
				s[0] = -1
			} else if hit == "last" {
				s[n-1] = -1
			}
			for _, impl := range []struct {
				name string
				fn   func([]T, T) bool
			}{
				{"archsimd", contains},
				{"slices", slices.Contains[[]T, T]},
			} {
				b.Run(fmt.Sprintf("%d/%s/%s", n, hit, impl.name), func(b *testing.B) {
					b.ReportAllocs()
					for b.Loop() {
						impl.fn(s, -1)
					}
				})
			}
		}
	}
}
