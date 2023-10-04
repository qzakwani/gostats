package gostats

import (
	"math"
)

func Mean[Z ~[]N, N number](n Z) float64 {
	l := len(n)
	if l == 0 {
		return math.NaN()
	}
	var sum float64 = 0
	sum = Sum(n)
	return sum / float64(l)
}

func GeometricMean[Z ~[]N, N number](n Z) float64 {
	l := len(n)
	if l == 0 {
		return math.NaN()
	}
	var product N = 1
	for _, n := range n {
		product *= n
	}
	return math.Pow(float64(product), 1.0/float64(l))
}

func HarmonicMean[Z ~[]N, N number](n Z) float64 {
	l := len(n)
	if l == 0 {
		return math.NaN()
	}
	var sum float64 = 0.0
	for _, i := range n {
		if i == 0 {
			return math.NaN()
		}
		sum += 1.0 / float64(i)
	}
	return float64(l) / sum
}

func WeightedMean[Z, W ~[]N, N number](numbers Z, weights W) float64 {
	if l := len(numbers); l != len(weights) || l == 0 {
		return math.NaN()
	}
	var sum float64 = 0.0
	var sumWeights float64 = 0.0
	for i, n := range numbers {
		w := weights[i]
		if w < 0 {
			return math.NaN()
		}
		sum += float64(n) * float64(w)
		sumWeights += float64(w)
	}
	return sum / sumWeights
}

func RMS[Z ~[]N, N number](n Z) float64 {
	l := len(n)
	if l == 0 {
		return math.NaN()
	}
	var sum float64 = 0
	for _, n := range n {
		sum += float64(n * n)
	}
	return math.Sqrt(sum / float64(l))
}

// for optimization
func MeanF(n []float64) float64 {
	l := len(n)
	if l == 0 {
		return math.NaN()
	}
	var sum float64 = 0
	sum = SumF(n)
	return sum / float64(l)
}

func GeometricMeanF(n []float64) float64 {
	l := len(n)
	if l == 0 {
		return math.NaN()
	}
	var product float64 = 1
	for _, n := range n {
		product *= n
	}
	return math.Pow(product, 1.0/float64(l))
}

func HarmonicMeanF(n []float64) float64 {
	l := len(n)
	if l == 0 {
		return math.NaN()
	}
	var sum float64 = 0.0
	for _, i := range n {
		if i == 0 {
			return math.NaN()
		}
		sum += 1.0 / i
	}
	return float64(l) / sum
}

func WeightedMeanF(numbers []float64, weights []float64) float64 {
	if l := len(numbers); l != len(weights) || l == 0 {
		return math.NaN()
	}
	var sum float64 = 0.0
	var sumWeights float64 = 0.0
	for i, n := range numbers {
		w := weights[i]
		if w < 0 {
			return math.NaN()
		}
		sum += n * w
		sumWeights += w
	}
	return sum / sumWeights
}

func RMSF(n []float64) float64 {
	l := len(n)
	if l == 0 {
		return math.NaN()
	}
	var sum float64 = 0
	for _, n := range n {
		sum += n * n
	}
	return math.Sqrt(sum / float64(l))
}
