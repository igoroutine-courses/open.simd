//go:build goexperiment.simd && arm64

package contains

import (
	"simd/archsimd"
	"slices"
)

// SliceContainsInt8 compares 16 elements at a time using arm64 Neon.
func SliceContainsInt8(s []int8, target int8) bool {
	const lanes = 16
	needle := archsimd.BroadcastInt8x16(target)
	for len(s) >= lanes {
		matches := archsimd.LoadInt8x16(s).Equal(needle).ToInt8x16()
		// True lanes are -1, false lanes are 0. ReduceMax would miss
		// a match unless every lane matched; ReduceMin detects any match.
		if matches.ReduceMin() != 0 {
			return true
		}
		s = s[lanes:]
	}
	return slices.Contains(s, target)
}

func SliceContainsInt32(s []int32, target int32) bool {
	const lanes = 4
	needle := archsimd.BroadcastInt32x4(target)
	for len(s) >= lanes {
		matches := archsimd.LoadInt32x4(s).Equal(needle).ToInt32x4()
		if matches.ReduceMin() != 0 {
			return true
		}
		s = s[lanes:]
	}
	return slices.Contains(s, target)
}

func SliceContainsInt64(s []int64, target int64) bool {
	const lanes = 2
	needle := archsimd.BroadcastInt64x2(target)
	for len(s) >= lanes {
		matches := archsimd.LoadInt64x2(s).Equal(needle).ToInt64x2()
		// Neon has no corresponding ReduceMin for two int64 lanes.
		if matches.GetElem(0)|matches.GetElem(1) != 0 {
			return true
		}
		s = s[lanes:]
	}
	return slices.Contains(s, target)
}
