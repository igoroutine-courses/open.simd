//go:build goexperiment.simd

package main

import (
	"fmt"
	"runtime"
	"simd"
)

// sumPortable works for any vector width, including software emulation.
// Go 1.27's portable API has no horizontal sum, so store the lanes once.
func sumPortable(x simd.Float32s) float32 {
	lanes := make([]float32, x.Len())
	x.Store(lanes)
	var total float32
	for _, value := range lanes {
		total += value
	}
	return total
}

// sumSlice accumulates using portable simd, then uses an architecture-specific
// horizontal reduction. Floating-point additions are reordered, so the result
// need not be bit-for-bit equal to a left-to-right scalar sum.
func sumSlice(xs []float32) float32 {
	var acc simd.Float32s
	for len(xs) >= acc.Len() {
		acc = acc.Add(simd.LoadFloat32s(xs))
		xs = xs[acc.Len():]
	}

	total := sumArch(acc)
	// Never perform a full-vector load past the end of the input.
	for _, x := range xs {
		total += x
	}

	return total
}

func main() {
	xs := []float32{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
		17, 18, 19, // Exercise a tail with 4-, 8- and 16-lane vectors.
	}

	fmt.Printf("%s, %s, vector=%d bits, emulated=%t\n",
		runtime.Version(), runtime.GOARCH, simd.VectorBitSize(), simd.Emulated())
	fmt.Println("Sum 1..19:", sumSlice(xs)) // 190
}
