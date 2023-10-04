package gostats

import (
	"math"
	"slices"
)

func Max[Z ~[]N, N number](n Z) float64 {
	if len(n) == 0 {
		return math.NaN()
	}
	return float64(slices.Max(n))
}

func MaxF(n []float64) float64 {
	if len(n) == 0 {
		return math.NaN()
	}
	return slices.Max(n)
}