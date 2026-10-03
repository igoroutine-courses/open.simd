//go:build goexperiment.simd && amd64

package main

import (
	"simd"
	"simd/archsimd"
)

func sumArch(x simd.Float32s) float32 {
	// ToArch returns the representation actually selected by simd. Merely
	// compiling for amd64 does not mean that AVX2 or AVX512 is available.
	switch v := x.ToArch().(type) {
	case archsimd.Float32x16:
		// Combine both 256-bit halves before reducing inside 128-bit groups.
		half := v.GetLo().Add(v.GetHi())
		half = half.ConcatAddPairsGrouped(half)
		half = half.ConcatAddPairsGrouped(half)
		return half.GetLo().GetElem(0) + half.GetHi().GetElem(0)
	case archsimd.Float32x8:
		v = v.ConcatAddPairsGrouped(v)
		v = v.ConcatAddPairsGrouped(v)
		// A third pairwise add would double both partial sums.
		return v.GetLo().GetElem(0) + v.GetHi().GetElem(0)
	default:
		return sumPortable(x)
	}
}
