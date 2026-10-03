//go:build !goexperiment.simd || (!amd64 && !arm64)

package contains

import "slices"

func SliceContainsInt8(s []int8, target int8) bool {
	return slices.Contains(s, target)
}

func SliceContainsInt32(s []int32, target int32) bool {
	return slices.Contains(s, target)
}

func SliceContainsInt64(s []int64, target int64) bool {
	return slices.Contains(s, target)
}
