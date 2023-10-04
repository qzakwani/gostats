package gostats

import (
	"math"

	"github.com/qzakwani/goslices"
)


func Sum[Z ~[]N, N number](n Z) float64 {
	if len(n) == 0 {
		return math.NaN()
	}
	return float64(goslices.Sum(&n))
}

func SumF(n []float64) float64 {
	if len(n) == 0 {
		return math.NaN()
	}
	return goslices.Sum(&n)
}