package gostats



func Avg[Z ~[]N, N number](n Z) float64 {
	return Mean(n)
}


func AvgF(n []float64) float64 {
	return MeanF(n)
}