package gostats

import (
	"math"

	"github.com/qzakwani/goslices"
)

func Median[Z ~[]N, N number](n Z) float64 {
	l := len(n)
	if l == 0 {
		return math.NaN()
	}
	if l == 1 {
		return float64(n[0])
	}
	m := goslices.Sort(n)
	if l%2 == 0 {
		return float64(m[l/2-1]+m[l/2]) / 2.0
	}
	return float64(m[l/2])
}

// optimized version for []float64
func MedianF(n []float64) float64 {
	l := len(n)
	if l == 0 {
		return math.NaN()
	}
	if l == 1 {
		return float64(n[0])
	}
	m := goslices.Sort(n)
	if l%2 == 0 {
		return float64(m[l/2-1]+m[l/2]) / 2.0
	}
	return float64(m[l/2])
}
