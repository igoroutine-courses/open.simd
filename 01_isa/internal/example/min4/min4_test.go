//go:build model_test

package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMin4(t *testing.T) {
	tests := []struct {
		name     string
		first    [4]int32
		second   [4]int32
		expected [4]int32
	}{
		{
			name:     "first all smaller",
			first:    [4]int32{1, 2, 3, 4},
			second:   [4]int32{10, 20, 30, 40},
			expected: [4]int32{1, 2, 3, 4},
		},
		{
			name:     "second all smaller",
			first:    [4]int32{10, 20, 30, 40},
			second:   [4]int32{1, 2, 3, 4},
			expected: [4]int32{1, 2, 3, 4},
		},
		{
			name:     "mixed lanes",
			first:    [4]int32{10, 5, 30, -100},
			second:   [4]int32{1, 20, 30, -50},
			expected: [4]int32{1, 5, 30, -100},
		},
		{
			name:     "equal values",
			first:    [4]int32{1, 2, 3, 4},
			second:   [4]int32{1, 2, 3, 4},
			expected: [4]int32{1, 2, 3, 4},
		},
		{
			name:     "negative values",
			first:    [4]int32{-1, -20, -300, -4000},
			second:   [4]int32{-2, -10, -400, -1000},
			expected: [4]int32{-2, -20, -400, -4000},
		},
		{
			name:     "positive and negative",
			first:    [4]int32{-1, 100, -300, 0},
			second:   [4]int32{1, -100, 300, -1},
			expected: [4]int32{-1, -100, -300, -1},
		},
		{
			name:     "zeroes",
			first:    [4]int32{0, 0, 0, 0},
			second:   [4]int32{0, 0, 0, 0},
			expected: [4]int32{0, 0, 0, 0},
		},
		{
			name:     "zero against negative",
			first:    [4]int32{0, 0, 0, 0},
			second:   [4]int32{-1, -2, -3, -4},
			expected: [4]int32{-1, -2, -3, -4},
		},
		{
			name:     "zero against positive",
			first:    [4]int32{0, 0, 0, 0},
			second:   [4]int32{1, 2, 3, 4},
			expected: [4]int32{0, 0, 0, 0},
		},
		{
			name:     "int32 max",
			first:    [4]int32{math.MaxInt32, 1, math.MaxInt32, -1},
			second:   [4]int32{0, math.MaxInt32, math.MaxInt32 - 1, math.MaxInt32},
			expected: [4]int32{0, 1, math.MaxInt32 - 1, -1},
		},
		{
			name:     "int32 min",
			first:    [4]int32{math.MinInt32, -1, math.MinInt32, 0},
			second:   [4]int32{0, math.MinInt32, math.MinInt32 + 1, math.MinInt32},
			expected: [4]int32{math.MinInt32, math.MinInt32, math.MinInt32, math.MinInt32},
		},
		{
			name:     "min and max together",
			first:    [4]int32{math.MinInt32, math.MaxInt32, -123, 456},
			second:   [4]int32{math.MaxInt32, math.MinInt32, 123, -456},
			expected: [4]int32{math.MinInt32, math.MinInt32, -123, -456},
		},
		{
			name:     "lane independence",
			first:    [4]int32{100, -100, 50, -50},
			second:   [4]int32{-100, 100, -50, 50},
			expected: [4]int32{-100, -100, -50, -50},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var got [4]int32

			Min4(&tt.first, &tt.second, &got)

			require.Equalf(t, tt.expected, got,
				"Min4(%v, %v) = %v, expected %v",
				tt.first, tt.second, got, tt.expected)
		})
	}
}

func TestMin4DoesNotModifyInputs(t *testing.T) {
	t.Parallel()

	first := [4]int32{1, -2, 3, -4}
	second := [4]int32{-10, 20, -30, 40}

	firstBefore := first
	secondBefore := second

	var dst [4]int32
	Min4(&first, &second, &dst)

	require.Equalf(t, firstBefore, first,
		"first was modified: got %v, expected %v",
		first, firstBefore)

	require.Equalf(t, secondBefore, second,
		"second was modified: got %v, expected %v",
		second, secondBefore)
}

func TestMin4OverwritesDst(t *testing.T) {
	t.Parallel()

	first := [4]int32{1, 2, 3, 4}
	second := [4]int32{10, -20, 30, -40}
	dst := [4]int32{999, 999, 999, 999}

	expected := [4]int32{1, -20, 3, -40}

	Min4(&first, &second, &dst)

	require.Equalf(t, expected, dst,
		"dst = %v, expected %v",
		dst, expected)
}

func TestMin4DstSameAsFirst(t *testing.T) {
	first := [4]int32{1, 100, -30, 40}
	second := [4]int32{10, 20, -3, 4}

	expected := [4]int32{1, 20, -30, 4}

	Min4(&first, &second, &first)

	require.Equalf(t, expected, first,
		"first/dst = %v, expected %v",
		first, expected)
}

func TestMin4DstSameAsSecond(t *testing.T) {
	first := [4]int32{1, 100, -30, 40}
	second := [4]int32{10, 20, -3, 4}

	expected := [4]int32{1, 20, -30, 4}

	Min4(&first, &second, &second)

	require.Equalf(t, expected, second,
		"second/dst = %v, expected %v",
		second, expected)
}

func TestMin4AgainstReferenceImplementation(t *testing.T) {
	cases := [][4]int32{
		{0, 0, 0, 0},
		{1, 2, 3, 4},
		{-1, -2, -3, -4},
		{math.MinInt32, math.MaxInt32, 0, -1},
		{123456, -654321, 42, -42},
	}

	for _, first := range cases {
		for _, second := range cases {
			t.Run("", func(t *testing.T) {
				var expected [4]int32
				min4Reference(&first, &second, &expected)

				var got [4]int32
				Min4(&first, &second, &got)

				require.Equalf(t, expected, got,
					"Min4(%v, %v) = %v, expected %v",
					first, second, got, expected)
			})
		}
	}
}
