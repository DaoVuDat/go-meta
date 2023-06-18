package main

import (
	"math/rand"
)

/*
Utility functions
*/
func cumulativeSum(arr []float64) []float64 {
	result := make([]float64, len(arr))

	sum := 0.0
	for i, v := range arr {
		sum += v
		result[i] = sum
	}
	return result
}

func RWS(probabilities []float64) int {
	index := 0

	cumsum := cumulativeSum(probabilities)
	r := rand.Float64()

	for i, v := range cumsum {
		if r <= v {
			index = i
			break
		}
	}

	return index
}
