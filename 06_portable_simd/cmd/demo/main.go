//go:build goexperiment.simd

package main

import (
	"fmt"
	"runtime"
	"simd"

	contains "github.com/igoroutine-courses/open.simd/06_portable_simd"
)

func main() {
	data := make([]byte, 19)
	for i := range data {
		data[i] = 7
	}
	data[len(data)-1] = 42
	fmt.Printf("%s: vector=%d bits, emulated=%t\n",
		runtime.GOARCH, simd.VectorBitSize(), simd.Emulated())
	fmt.Println("19 bytes, target in last byte:", contains.SliceContains(data, 42))
	fmt.Println("Missing target:", contains.SliceContains(data, 99))
	fmt.Println("Partial load, target=0 (padding must not match):",
		contains.SliceContainsPart([]byte{1, 2, 3}, 0))
}
