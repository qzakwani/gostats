package gostats

import "math"


func SD[Z ~[]N, N number](n Z) float64 {
	return math.Sqrt(Var(n))
}

func SDF(n []float64) float64 {
	return math.Sqrt(VarF(n))
}