package main

import (
	"math"
)

func main() {

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

	// Setup objective function
	objectiveFunction := func(position []float64) float64 {
		sum := 0.0
		for _, pos := range position {
			sum += math.Pow(pos, 2)
		}
		return sum
	}

	// create gwo struct
	gwo := Gwo{
		numberOfWolf:      30,
		dimension:         dimension,
		iteration:         500,
		lowerBound:        lowerBound,
		upperBound:        upperBound,
		objectiveFunction: objectiveFunction,
	}

	gwo.Run()

	//for i, v := range gwo.population {
	//	fmt.Printf("%d - %v\n", i, &v)
	//}
}
