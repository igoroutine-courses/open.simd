package main

import "fmt"

func main() {
	first := [4]int32{1, 3, 5, 7}
	second := [4]int32{2, 4, 6, 8}
	dst := [4]int32{}

	Min4(&first, &second, &dst)
	fmt.Println(dst) // [1 3 5 7]
}

// Min4 computes the element-wise signed minimum of two vectors of four int32 values
// and stores the result in dst.
func Min4(first, second, dst *[4]int32)
