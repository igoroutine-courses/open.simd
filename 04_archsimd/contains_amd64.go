//go:build goexperiment.simd && amd64

package contains

import (
	"simd/archsimd"
	"slices"
)

func SliceContainsInt8(s []int8, target int8) bool {
	if !archsimd.X86.AVX2() {
		return slices.Contains(s, target)
	}

	const lanes = 32

	i := 0
	n := len(s) - len(s)%lanes
	needle := archsimd.BroadcastInt8x32(target)

	for ; i < n; i += lanes {
		v := archsimd.LoadInt8x32(s[i:])
		if v.Equal(needle).ToBits() != 0 {
			return true
		}
	}

	for ; i < len(s); i++ {
		if s[i] == target {
			return true
		}
	}

	return false
}

func SliceContainsInt32(s []int32, target int32) bool {
	if !archsimd.X86.AVX2() {
		return slices.Contains(s, target)
	}

	const lanes = 8

	i := 0
	n := len(s) - len(s)%lanes
	needle := archsimd.BroadcastInt32x8(target)

	for ; i < n; i += lanes {
		v := archsimd.LoadInt32x8(s[i:])
		if v.Equal(needle).ToBits() != 0 {
			return true
		}
	}

	for ; i < len(s); i++ {
		if s[i] == target {
			return true
		}
	}

	return false
}

func SliceContainsInt64(s []int64, target int64) bool {
	if !archsimd.X86.AVX2() {
		return slices.Contains(s, target)
	}

	const lanes = 4

	i := 0
	n := len(s) - len(s)%lanes
	needle := archsimd.BroadcastInt64x4(target)

	for ; i < n; i += lanes {
		v := archsimd.LoadInt64x4(s[i:])
		if v.Equal(needle).ToBits() != 0 {
			return true
		}
	}

	for ; i < len(s); i++ {
		if s[i] == target {
			return true
		}
	}

	return false
}
