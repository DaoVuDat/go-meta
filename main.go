package main

import (
	"fmt"
	"math"
	"time"
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

	start := time.Now()
	gwo.Run()
	end := time.Since(start)
	fmt.Printf("GWO tooks %s\n", end)
	fmt.Printf("==> Best value: %E\n", gwo.alpha.value)

	avoa := Avoa{
		numberOfWolf:      30,
		dimension:         dimension,
		iteration:         500,
		lowerBound:        lowerBound,
		upperBound:        upperBound,
		objectiveFunction: objectiveFunction,
		p1:                0.6,
		p2:                0.4,
		p3:                0.6,
		alpha:             0.8,
		betha:             0.2,
		gamma:             2.5,
	}

	start = time.Now()
	avoa.Run()
	end = time.Since(start)
	fmt.Printf("AVOA tooks %s\n", end)
	fmt.Printf("==> Best value: %E\n", avoa.best1.value)

}
