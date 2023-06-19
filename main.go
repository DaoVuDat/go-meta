package main

import (
	"fmt"
	"time"
)

func main() {

	dimension, lowerBound, upperBound, objectiveProblem := GetFunction(9)

	// create gwo struct
	gwo := Gwo{
		numberOfWolf:      30,
		dimension:         dimension,
		iteration:         500,
		lowerBound:        lowerBound,
		upperBound:        upperBound,
		objectiveFunction: objectiveProblem,
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
		objectiveFunction: objectiveProblem,
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
