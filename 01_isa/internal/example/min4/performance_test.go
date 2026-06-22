//go:build performance_test

package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMin4SIMDPerformance(t *testing.T) {
	const targetSize = 64 << 20

	simd := testing.Benchmark(func(b *testing.B) {
		first := filledSlice[int32](math.MinInt32, targetSize)
		second := filledSlice[int32](math.MaxInt32, targetSize)
		dst := filledSlice[int32](0, targetSize)

		for b.Loop() {
			for i := 0; i < targetSize/4; i++ {
				f := (*[4]int32)(first[i*4 : (i+1)*4])
				s := (*[4]int32)(second[i*4 : (i+1)*4])
				d := (*[4]int32)(dst[i*4 : (i+1)*4])

				Min4(f, s, d)
			}
		}

		require.Equal(t, first, dst)
	})

	reference := testing.Benchmark(func(b *testing.B) {
		first := filledSlice[int32](math.MinInt32, targetSize)
		second := filledSlice[int32](math.MaxInt32, targetSize)
		dst := filledSlice[int32](0, targetSize)

		for b.Loop() {
			for i := 0; i < targetSize/4; i++ {
				f := (*[4]int32)(first[i*4 : (i+1)*4])
				s := (*[4]int32)(second[i*4 : (i+1)*4])
				d := (*[4]int32)(dst[i*4 : (i+1)*4])
				min4Reference(f, s, d)
			}
		}

		require.Equal(t, first, dst)
	})

	ratio := float64(simd.NsPerOp()) / float64(reference.NsPerOp())
	speedup := 1 / ratio

	t.Logf("simd:      %d", simd.NsPerOp())
	t.Logf("reference: %d", reference.NsPerOp())
	t.Logf("ratio simd/reference: %.3f", ratio)
	t.Logf("speedup: %.2fx", speedup)

	require.GreaterOrEqual(t, speedup, 1.8)
}
