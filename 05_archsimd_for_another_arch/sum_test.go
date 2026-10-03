//go:build goexperiment.simd

package main

import (
	"math"
	"simd"
	"testing"
)

func TestSumWidthsAndTails(t *testing.T) {
	width := (simd.Float32s{}).Len()
	t.Logf("vector=%d bits, emulated=%t", simd.VectorBitSize(), simd.Emulated())
	for n := 0; n <= 4*width+3; n++ {
		backing := make([]float32, n+2)
		backing[0], backing[n+1] = 1e9, 1e9
		xs := backing[1 : n+1 : n+1]
		var want float32
		for i := range xs {
			xs[i] = float32(i%17 - 8)
			want += xs[i] // Small integers: addition is exact in float32.
		}
		if got := sumSlice(xs); got != want {
			t.Fatalf("n=%d width=%d: got %v, want %v", n, width, got, want)
		}
		for i, value := range xs {
			if value != float32(i%17-8) {
				t.Fatalf("input modified at %d", i)
			}
		}
	}

	// Every lane contributes exactly once; catches the old extra pairwise add.
	v := simd.BroadcastFloat32s(1)
	if got := sumArch(v); got != float32(width) {
		t.Fatalf("horizontal sum: got %v, want %d", got, width)
	}
	if got := sumPortable(v); got != float32(width) {
		t.Fatalf("portable sum: got %v, want %d", got, width)
	}
}

func TestSumRounding(t *testing.T) {
	xs := make([]float32, 1003)
	var want, magnitude float64
	for i := range xs {
		xs[i] = float32(math.Sin(float64(i)) / 3)
		want += float64(xs[i])
		magnitude += math.Abs(float64(xs[i]))
	}
	got := float64(sumSlice(xs))
	// Different reduction orders need not agree bit for bit.
	if math.Abs(got-want) > 1e-5*magnitude {
		t.Fatalf("got %v, want approximately %v", got, want)
	}
}

func TestSumSpecialValues(t *testing.T) {
	width := (simd.Float32s{}).Len()
	for _, special := range []float32{float32(math.Inf(1)), float32(math.Inf(-1)), float32(math.NaN())} {
		for _, pos := range []int{0, width} {
			xs := make([]float32, width+1)
			xs[pos] = special
			got := sumSlice(xs)
			if math.IsNaN(float64(special)) {
				if !math.IsNaN(float64(got)) {
					t.Fatalf("expected NaN, got %v", got)
				}
			} else if got != special {
				t.Fatalf("got %v, want %v", got, special)
			}
		}
	}
}
