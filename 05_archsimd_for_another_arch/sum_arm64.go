//go:build goexperiment.simd && arm64

package main

import (
	"simd"
	"simd/archsimd"
)

func sumArch(x simd.Float32s) float32 {
	if v, ok := x.ToArch().(archsimd.Float32x4); ok {
		v = v.ConcatAddPairs(v) // [a+b, c+d, a+b, c+d]
		v = v.ConcatAddPairs(v) // Each lane contains a+b+c+d.
		return v.GetElem(0)
	}
	return sumPortable(x)
}
