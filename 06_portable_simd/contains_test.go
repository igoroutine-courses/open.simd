//go:build goexperiment.simd

package contains

import (
	"fmt"
	"simd"
	"slices"
	"testing"
)

func TestContains(t *testing.T) {
	width := (simd.Uint8s{}).Len()
	t.Logf("vector=%d bits, emulated=%t", simd.VectorBitSize(), simd.Emulated())
	for _, impl := range []struct {
		name string
		fn   func([]byte, byte) bool
	}{{"full", SliceContains}, {"partial", SliceContainsPart}} {
		t.Run(impl.name, func(t *testing.T) {
			if impl.fn(nil, 0) {
				t.Fatal("nil input must not contain zero")
			}
			for n := 0; n <= 2*width+3; n++ {
				for _, target := range []byte{0, 1, 42, 128, 255} {
					backing := make([]byte, n+2)
					backing[0], backing[n+1] = target, target
					s := backing[1 : n+1 : n+1]
					for i := range s {
						s[i] = target + 1
					}
					if impl.fn(s, target) {
						t.Fatalf("matched padding/outside input: n=%d target=%d", n, target)
					}
					for pos := range s {
						s[pos] = target
						before := slices.Clone(backing)
						if got := impl.fn(s, target); got != slices.Contains(s, target) {
							t.Fatalf("n=%d pos=%d target=%d: got %t", n, pos, target, got)
						}
						if !slices.Equal(backing, before) {
							t.Fatal("input was modified")
						}
						s[pos] = target + 1
					}
				}
			}
		})
	}
}

func FuzzContains(f *testing.F) {
	width := (simd.Uint8s{}).Len()
	for _, n := range []int{0, 1, width - 1, width, width + 1, 2*width + 3} {
		s := make([]byte, n)
		for i := range s {
			s[i] = 7
		}
		f.Add(s, byte(0)) // Zero-filled inactive lanes must not match.
		if n > 0 {
			s[n-1] = 42
		}
		f.Add(s, byte(42))
	}
	f.Fuzz(func(t *testing.T, s []byte, target byte) {
		before := slices.Clone(s)
		want := slices.Contains(s, target)
		if got := SliceContains(s, target); got != want {
			t.Fatalf("full: got %t, want %t", got, want)
		}
		if got := SliceContainsPart(s, target); got != want {
			t.Fatalf("partial: got %t, want %t", got, want)
		}
		if !slices.Equal(s, before) {
			t.Fatal("input modified")
		}
	})
}

func BenchmarkContains(b *testing.B) {
	for _, n := range []int{19, 1024, 100_000} {
		for _, hit := range []string{"missing", "first", "last"} {
			s := make([]byte, n)
			if hit == "first" {
				s[0] = 42
			} else if hit == "last" {
				s[n-1] = 42
			}
			for _, impl := range []struct {
				name string
				fn   func([]byte, byte) bool
			}{
				{"simd", SliceContains},
				{"simd-part", SliceContainsPart},
				{"slices", slices.Contains[[]byte, byte]},
			} {
				b.Run(fmt.Sprintf("%d/%s/%s", n, hit, impl.name), func(b *testing.B) {
					b.ReportAllocs()
					for b.Loop() {
						impl.fn(s, 42)
					}
				})
			}
		}
	}
}
