package gostats

import "testing"


var testSlice = []float64{1,2,3,4,5,6,7,8,9,10}

func BenchmarkVar(b *testing.B) {
	for	i := 0; i < b.N; i++ {
		Var(testSlice)
	}
}

func BenchmarkVarF(b *testing.B) {
	for	i := 0; i < b.N; i++ {
		VarF(testSlice)
	}
}