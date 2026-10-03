//go:build goexperiment.simd && !amd64 && !arm64

package main

import "simd"

func sumArch(x simd.Float32s) float32 {
	return sumPortable(x)
}
