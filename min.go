package gostats

import (
	"math"
	"slices"
)

func Min[Z ~[]N, N number](n Z) float64 {
	if len(n) == 0 {
		return math.NaN()
	}
	return float64(slices.Min(n))
}

func MinF(n []float64) float64 {
	if len(n) == 0 {
		return math.NaN()
	}
	return slices.Min(n)
}