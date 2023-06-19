package main

import (
	"math"
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

func LevyFlight(dimension int) []float64 {
	beta := 1.5
	sigma := math.Pow((math.Gamma(1.0+beta)*math.Sin(math.Pi*beta/2.0))/
		(math.Gamma((1+beta)/2)*beta*math.Pow(2, (beta-1)/2)), 1/beta)

	o := make([]float64, dimension)
	for i := range o {
		u := rand.Float64() * sigma
		v := rand.Float64()
		step := u / (math.Pow(math.Abs(v), 1/beta))
		o[i] = step
	}

	return o
}
