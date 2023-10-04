package gostats

import (
	"math"
)

func Mode[Z ~[]N, N number](n Z) float64 {
	if len(n) == 0 {
		return math.NaN()
	}

	// c:= goslices.NewCounter(n)
	

	return 0
}