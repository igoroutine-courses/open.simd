package contains

import (
	"math"
	stdslices "slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSliceContainsInt8(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		s      []int8
		target int8
	}{
		{"empty", []int8{}, 1},
		{"single found", []int8{7}, 7},
		{"single not found", []int8{7}, 8},
		{"first", []int8{9, 1, 2, 3}, 9},
		{"middle", []int8{1, 2, -5, 3}, -5},
		{"last", []int8{1, 2, 3, 4}, 4},
		{"tail", append(make([]int8, 37), -42), -42},
		{"min", []int8{math.MinInt8, 0, 1}, math.MinInt8},
		{"max", []int8{0, 1, math.MaxInt8}, math.MaxInt8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, stdslices.Contains(tt.s, tt.target), SliceContainsInt8(tt.s, tt.target))
		})
	}
}

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
		{"tail", append(make([]int32, 19), -42), -42},
		{"min", []int32{math.MinInt32, 0, 1}, math.MinInt32},
		{"max", []int32{0, 1, math.MaxInt32}, math.MaxInt32},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, stdslices.Contains(tt.s, tt.target), SliceContainsInt32(tt.s, tt.target))
		})
	}
}

func TestSliceContainsInt64(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		s      []int64
		target int64
	}{
		{"empty", []int64{}, 1},
		{"single found", []int64{7}, 7},
		{"single not found", []int64{7}, 8},
		{"first", []int64{9, 1, 2, 3}, 9},
		{"middle", []int64{1, 2, -5, 3}, -5},
		{"last", []int64{1, 2, 3, 4}, 4},
		{"tail", append(make([]int64, 11), -42), -42},
		{"min", []int64{math.MinInt64, 0, 1}, math.MinInt64},
		{"max", []int64{0, 1, math.MaxInt64}, math.MaxInt64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, stdslices.Contains(tt.s, tt.target), SliceContainsInt64(tt.s, tt.target))
		})
	}
}
