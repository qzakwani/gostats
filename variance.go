package gostats

import (
	"math"
)


func Var[Z ~[]N, N number](n Z) float64 {
	l:=len(n)
	if l == 0 {
		return math.NaN()
	}

	m:=Mean(n)
	var sum float64
	for _, v := range n {
		i:=float64(v)
		sum += (i - m) * (i - m)
	}
	return sum /float64(l)
}


func VarF(n []float64) float64 {
	l:=len(n)
	if l == 0 {
		return math.NaN()
	}

	m:=MeanF(n)
	var sum float64
	for _, v := range n {
		sum += (v - m) * (v - m)
	}
	return sum /float64(l)
}