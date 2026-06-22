//go:build model_test

package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSub(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		first     int64
		second    int64
		expResult int64
	}{
		{
			name:      "small positive",
			first:     1,
			second:    2,
			expResult: -1,
		},
		{
			name:      "small negative",
			first:     -1,
			second:    -2,
			expResult: 1,
		},
		{
			name:      "mixed signs",
			first:     -3,
			second:    2,
			expResult: -5,
		},
		{
			name:      "positive minus negative",
			first:     10,
			second:    -5,
			expResult: 15,
		},
		{
			name:      "zero minus positive",
			first:     0,
			second:    7,
			expResult: -7,
		},
		{
			name:      "zero minus negative",
			first:     0,
			second:    -7,
			expResult: 7,
		},
		{
			name:      "same values",
			first:     42,
			second:    42,
			expResult: 0,
		},
		{
			name:      "big positive",
			first:     math.MaxInt32,
			second:    math.MaxInt32 - 1,
			expResult: 1,
		},
		{
			name:      "big negative",
			first:     math.MinInt32 + 1,
			second:    math.MinInt32,
			expResult: 1,
		},
		{
			name:      "max int64 minus one",
			first:     math.MaxInt64,
			second:    1,
			expResult: math.MaxInt64 - 1,
		},
		{
			name:      "min int64 plus one",
			first:     math.MinInt64 + 1,
			second:    1,
			expResult: math.MinInt64,
		},
		{
			name:      "negative minus max int32",
			first:     -1,
			second:    math.MaxInt32,
			expResult: -1 - math.MaxInt32,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.expResult, Sub(tt.first, tt.second))
		})
	}
}
