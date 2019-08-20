package core

import (
	"fmt"
	"math/rand"
)

func Debugln(a ...interface{}) {
	if DEBUG {
		fmt.Println(a...)
	}
}
func Debugf(s string, a ...interface{}) {
	if DEBUG {
		fmt.Printf(s, a...)
	}
}

func NormalizedRandomArray(min, sum float64, size int) []float64 {
	// This fucntion creates random array of size size with min min and
	// arr sums to sum
	s, arr := 0.0, make([]float64, size)
	for i := range arr {
		arr[i] = min + rand.Float64()
		s += arr[i]
	}
	// normalize to sum
	for i := range arr {
		arr[i] = arr[i] * sum / s
	}
	return arr
}
