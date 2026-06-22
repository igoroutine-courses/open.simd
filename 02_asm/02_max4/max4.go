package main

import "fmt"

func main() {
	first := [4]int32{1, 2, 3, 4}
	second := [4]int32{2, 4, 6, 8}
	dst := [4]int32{}

	Max4(&first, &second, &dst)
	fmt.Println(dst)
}

// Max4 computes the element-wise signed maximum of two vectors of four int32 values
// and stores the result in dst.
func Max4(first, second, dst *[4]int32)
