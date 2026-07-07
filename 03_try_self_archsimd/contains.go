package main

import "fmt"

func main() {
	s := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	fmt.Println("Contains 7:", SliceContainsInt32(s, 7))
	fmt.Println("Contains 99:", SliceContainsInt32(s, 99))
}

func SliceContainsInt32(s []int32, target int32) bool

type Int32x8 struct {
	values [8]int32
}

type MaskInt32x8 struct {
	values [8]int32
}

func BroadcastInt32x8(value int32) Int32x8
func LoadInt32x8(ptr *int32) Int32x8
func EqualInt32x8(a, b Int32x8) MaskInt32x8
func MaskToBits(mask MaskInt32x8) uint32
