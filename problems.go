package main

import (
	"fmt"
	"math"
)

func GetFunction(problemNumber int) (
	dimension int,
	lowerBound []float64,
	upperBound []float64,
	objectiveFunction func([]float64) float64) {
	switch problemNumber {
	case 1:
		// Sphere
		dimension := 30

		lowerValue := -100.0
		upperValue := 100.0
		lowerBound := make([]float64, dimension)
		upperBound := make([]float64, dimension)

		// Initial lowerBound and upperBound - for the same values
		for i := 0; i < dimension; i++ {
			upperBound[i] = upperValue
			lowerBound[i] = lowerValue
		}
		return dimension, lowerBound, upperBound, sphere
	case 5:
		// F5
		dimension := 30

		lowerValue := -30.0
		upperValue := 30.0
		lowerBound := make([]float64, dimension)
		upperBound := make([]float64, dimension)

		// Initial lowerBound and upperBound - for the same values
		for i := 0; i < dimension; i++ {
			upperBound[i] = upperValue
			lowerBound[i] = lowerValue
		}
		return dimension, lowerBound, upperBound, f5
	case 9:
		// F9
		dimension := 30

		lowerValue := -5.12
		upperValue := 5.12
		lowerBound := make([]float64, dimension)
		upperBound := make([]float64, dimension)

		// Initial lowerBound and upperBound - for the same values
		for i := 0; i < dimension; i++ {
			upperBound[i] = upperValue
			lowerBound[i] = lowerValue
		}
		return dimension, lowerBound, upperBound, f9
	default:
		// Return problem 1 (sphere)
		fmt.Printf("Do not have F%d => Use F1 instead\n", problemNumber)
		dimension := 30

		lowerValue := -100.0
		upperValue := 100.0
		lowerBound := make([]float64, dimension)
		upperBound := make([]float64, dimension)

		// Initial lowerBound and upperBound - for the same values
		for i := 0; i < dimension; i++ {
			upperBound[i] = upperValue
			lowerBound[i] = lowerValue
		}
		return dimension, lowerBound, upperBound, sphere
	}
}

// Objective Functions

func sphere(x []float64) float64 {
	sum := 0.0
	for _, pos := range x {
		sum += math.Pow(pos, 2)
	}
	return sum
}

func f5(x []float64) float64 {
	dimension := len(x)
	sum := 0.0
	for i := 0; i < dimension-1; i++ {
		sum += 100*math.Pow(x[i+1]-x[i]*x[i], 2) + math.Pow(x[i]-1, 2)
	}
	return sum
}

func f9(x []float64) float64 {
	dimension := float64(len(x))
	sum := 0.0
	for _, v := range x {
		sum += v*v - 10*math.Cos(2*math.Pi*v)
	}
	return 10*dimension + sum
}
