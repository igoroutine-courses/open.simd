package main

import (
	"math"
	stdslices "slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSliceContainsInt32(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		s      []int32
		target int32
	}{
		{"empty", []int32{}, 1},
		{"single found", []int32{7}, 7},
		{"single not found", []int32{7}, 8},
		{"first", []int32{9, 1, 2, 3}, 9},
		{"middle", []int32{1, 2, -5, 3}, -5},
		{"last", []int32{1, 2, 3, 4}, 4},
		{"full simd block", []int32{1, 2, 3, 4}, 4},
		{"simd + tail", []int32{1, 2, 3, 4, 5, 6}, 6},
		{"tail only", []int32{1, 2, 3}, 2},
		{"two simd blocks", []int32{1, 2, 3, 4, 5, 6, 7, 8}, 7},
		{"not found large", append(make([]int32, 100), 42), 99},
		{"found at end large", append(make([]int32, 100), 42), 42},
		{"min", []int32{math.MinInt32, 0, 1}, math.MinInt32},
		{"max", []int32{0, 1, math.MaxInt32}, math.MaxInt32},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			expected := stdslices.Contains(tt.s, tt.target)
			actual := SliceContainsInt32(tt.s, tt.target)
			require.Equal(t, expected, actual, "slice=%v, target=%d", tt.s, tt.target)
		})
	}
}

func TestSliceContainsInt32Performance(t *testing.T) {
	const size = 100_000

	s := make([]int32, size)
	for i := range s {
		s[i] = int32(i + 1)
	}
	target := int32(-1) // Not in slice - worst case

	solution := testing.Benchmark(func(b *testing.B) {
		for b.Loop() {
			SliceContainsInt32(s, target)
		}
	})

	reference := testing.Benchmark(func(b *testing.B) {
		for b.Loop() {
			stdslices.Contains(s, target)
		}
	})

	speedup := float64(reference.NsPerOp()) / float64(solution.NsPerOp())

	t.Logf("SIMD solution:     %d ns/op", solution.NsPerOp())
	t.Logf("Scalar reference:  %d ns/op", reference.NsPerOp())
	t.Logf("Speedup: %.2fx", speedup)

	require.GreaterOrEqual(t, speedup, 1.5, "SIMD should be at least 1.5x faster")
}
