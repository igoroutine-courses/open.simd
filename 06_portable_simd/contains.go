//go:build goexperiment.simd

// Package contains demonstrates Go 1.27's architecture-independent SIMD API.
package contains

import (
	"simd"
	"slices"
)

// SliceContains reports whether target occurs in s, without assuming a vector
// width or importing archsimd. The portable mask API has no Any/ToBits in Go
// 1.27, so we accumulate matches and reduce the stored lanes at the end.
// Unlike slices.Contains, this variant scans every full block, even on a hit.
func SliceContains(s []byte, target byte) bool {
	needle := simd.BroadcastUint8s(target)
	width := needle.Len()
	if len(s) < width {
		return slices.Contains(s, target)
	}

	var matches simd.Mask8s
	for len(s) >= width {
		matches = matches.Or(simd.LoadUint8s(s).Equal(needle))
		s = s[width:]
	}
	if slices.Contains(s, target) {
		return true
	}

	lanes := make([]int8, width)
	matches.ToInt8s().Store(lanes)
	for _, lane := range lanes {
		if lane != 0 {
			return true
		}
	}
	return false
}

// SliceContainsPart demonstrates partial loads and stores, including the final
// short block. This is a teaching example, not an assertion of better speed.
func SliceContainsPart(s []byte, target byte) bool {
	if len(s) == 0 {
		return false
	}
	needle := simd.BroadcastUint8s(target)
	lanes := make([]int8, needle.Len())
	for len(s) > 0 {
		v, loaded := simd.LoadUint8sPart(s)
		// A partial load zero-fills inactive lanes. If target == 0, those
		// lanes compare equal too! Store and inspect only actual input lanes.
		v.Equal(needle).ToInt8s().StorePart(lanes[:loaded])
		for _, lane := range lanes[:loaded] {
			if lane != 0 {
				return true
			}
		}
		s = s[loaded:]
	}
	return false
}
