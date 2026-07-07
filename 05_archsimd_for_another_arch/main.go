//go:build goexperiment.simd

package main

import (
	"fmt"

	"simd"
	"simd/archsimd"
)

func sumPortable(x simd.Float32s) float32 {
	switch a := x.ToArch().(type) {
	case archsimd.Float32x16:
		a = a.AddPairsGrouped(a)
		a = a.AddPairsGrouped(a)
		a = a.AddPairsGrouped(a)
		a = a.AddPairsGrouped(a)

		return a.GetLo().GetElem(0)

	case archsimd.Float32x8:
		a = a.AddPairsGrouped(a) // 8 -> 4
		a = a.AddPairsGrouped(a) // 4 -> 2
		a = a.AddPairsGrouped(a) // 2 -> 1

		return a.GetLo().GetElem(0) + a.GetHi().GetElem(0)

	case archsimd.Float32x4:
		a = a.AddPairsGrouped(a) // 4 -> 2
		a = a.AddPairsGrouped(a) // 2 -> 1

		return a.GetElem(0)

	default:

		var tmp [64]float32
		x.StoreSlice(tmp[:x.Len()])

		var sum float32
		for i := 0; i < x.Len(); i++ {
			sum += tmp[i]
		}
		return sum
	}
}

func sumSlice(xs []float32) float32 {
	var total float32

	for len(xs) >= simd.Float32s{}.Len() {
		v := simd.LoadFloat32Slice(xs)
		total += sumPortable(v)
		xs = xs[v.Len():]
	}

	// scalar tail
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
	}

	fmt.Println(sumSlice(xs))
}
