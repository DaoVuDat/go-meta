package main

import (
	"fmt"
	"math"
	"math/rand"
)

type ObjectiveFunction func([]float64) float64

type wolf struct {
	positions []float64
	value     float64
}

type Gwo struct {
	lowerBound        []float64
	upperBound        []float64
	dimension         int
	numberOfWolf      int
	iteration         int
	population        []*wolf // this must be a slice of pointer because of the assignment in for range
	objectiveFunction ObjectiveFunction
	alpha             wolf
	beta              wolf
	delta             wolf
}

func (gwo *Gwo) Run() {
	// Initialize
	gwo.initialization()
	gwo.findBest()

	// Loop algorithm
	for currentIteration := 0; currentIteration < gwo.iteration; currentIteration++ {
		// A decrease linearly from 2 to 0
		a := 2 - float64(currentIteration)*(2.0/float64(gwo.iteration))
		// Update position of each wolf
		for _, wolf := range gwo.population {
			for iPos, pos := range wolf.positions {
				r1 := rand.Float64()
				r2 := rand.Float64()

				A1 := 2.0*a*r1 - a
				C1 := 2.0 * r2

				DAlpha := math.Abs(C1*gwo.alpha.positions[iPos] - pos)
				X1 := gwo.alpha.positions[iPos] - A1*DAlpha

				r1 = rand.Float64()
				r2 = rand.Float64()

				A2 := 2.0*a*r1 - a
				C2 := 2.0 * r2

				DBeta := math.Abs(C2*gwo.beta.positions[iPos] - pos)
				X2 := gwo.beta.positions[iPos] - A2*DBeta

				r1 = rand.Float64()
				r2 = rand.Float64()

				A3 := 2.0*a*r1 - a
				C3 := 2.0 * r2

				DDelta := math.Abs(C3*gwo.delta.positions[iPos] - pos)
				X3 := gwo.delta.positions[iPos] - A3*DDelta

				newPos := (X1 + X2 + X3) / 3

				if newPos > gwo.upperBound[iPos] {
					newPos = gwo.upperBound[iPos]
				}

				if newPos < gwo.lowerBound[iPos] {
					newPos = gwo.lowerBound[iPos]
				}

				wolf.positions[iPos] = newPos
			}
			wolf.value = gwo.objectiveFunction(wolf.positions)
			//fmt.Printf("%d - %p - %v\n", currentIteration, wolf, wolf.positions)
		}

		// Find top 3 wolves
		gwo.findBest()

		// Print information
		fmt.Printf("Iteration %d ", currentIteration+1)
		fmt.Printf("=> Best value: %E\n", gwo.alpha.value)
	}
}

func (gwo *Gwo) findBest() {
	for _, wolf := range gwo.population {
		if wolf.value < gwo.alpha.value {
			gwo.alpha.value = wolf.value
			copy(gwo.alpha.positions, wolf.positions)
		} else if wolf.value < gwo.beta.value {
			gwo.beta.value = wolf.value
			copy(gwo.beta.positions, wolf.positions)
		} else if wolf.value < gwo.delta.value {
			gwo.delta.value = wolf.value
			copy(gwo.delta.positions, wolf.positions)
		}
	}
}

func (gwo *Gwo) initialization() {
	for i := 0; i < gwo.numberOfWolf; i++ {
		newPosition := make([]float64, gwo.dimension)
		for j := 0; j < gwo.dimension; j++ {
			newPosDim := gwo.lowerBound[j] + rand.Float64()*(gwo.upperBound[j]-gwo.lowerBound[j])
			newPosition[j] = newPosDim
		}

		value := gwo.objectiveFunction(newPosition)
		gwo.population = append(gwo.population, &wolf{positions: newPosition, value: value})
	}

	gwo.alpha = wolf{
		value:     math.MaxFloat64,
		positions: make([]float64, gwo.dimension),
	}

	gwo.beta = wolf{
		value:     math.MaxFloat64,
		positions: make([]float64, gwo.dimension),
	}

	gwo.delta = wolf{
		value:     math.MaxFloat64,
		positions: make([]float64, gwo.dimension),
	}
}
