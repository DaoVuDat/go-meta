package main

import (
	"fmt"
	"math"
	"math/rand"
)

type vulture struct {
	positions []float64
	value     float64
}

type Avoa struct {
	lowerBound        []float64
	upperBound        []float64
	dimension         int
	numberOfWolf      int
	iteration         int
	population        []vulture
	objectiveFunction ObjectiveFunction
	best1             vulture
	best2             vulture
	p1                float64
	p2                float64
	p3                float64
	alpha             float64
	betha             float64
	gamma             float64
}

func (avoa *Avoa) initialization() {
	for i := 0; i < avoa.numberOfWolf; i++ {
		newPosition := make([]float64, avoa.dimension)
		for j := 0; j < avoa.dimension; j++ {
			newPosDim := avoa.lowerBound[j] + rand.Float64()*(avoa.upperBound[j]-avoa.lowerBound[j])
			newPosition[j] = newPosDim
		}

		value := avoa.objectiveFunction(newPosition)
		avoa.population = append(avoa.population, vulture{positions: newPosition, value: value})
	}

	avoa.best1 = vulture{
		value:     math.MaxFloat64,
		positions: make([]float64, avoa.dimension),
	}

	avoa.best2 = vulture{
		value:     math.MaxFloat64,
		positions: make([]float64, avoa.dimension),
	}

}

func (avoa *Avoa) findBest() {
	for i, _ := range avoa.population {
		if avoa.population[i].value < avoa.best1.value {
			avoa.best1.value = avoa.population[i].value
			copy(avoa.best1.positions, avoa.population[i].positions)
		} else if avoa.population[i].value < avoa.best2.value {
			avoa.best2.value = avoa.population[i].value
			copy(avoa.best2.positions, avoa.population[i].positions)
		}
	}
}

func (avoa *Avoa) Run() {
	// Initialize
	avoa.initialization()
	avoa.findBest()

	// Loop algorithm
	for currentIteration := 0; currentIteration < avoa.iteration; currentIteration++ {
		random2ToMinus2 := rand.Float64()*(2.0+2.0) - 2

		a := random2ToMinus2 * (math.Pow(math.Sin((math.Pi/2)*(float64(currentIteration)/float64(avoa.iteration))), avoa.gamma) +
			math.Cos((math.Pi/2)*(float64(currentIteration)/float64(avoa.iteration))) - 1)

		P1 := (2*rand.Float64()+1)*(1-(float64(currentIteration)/float64(avoa.iteration))) + a

		// Update position of each wolf
		for i, _ := range avoa.population {
			vult := &avoa.population[i]
			F := P1 * (2.0*rand.Float64() - 1)

			// random select vulture
			randomSelectVultureIndex := RWS([]float64{avoa.alpha, avoa.betha})

			if math.Abs(F) >= 1 {
				// Exploration
				avoa.exploration(randomSelectVultureIndex, F, vult)
			} else {
				// Exploitation
			}
		}

		// Find top 2 vultures
		avoa.findBest()

		// Print information
		fmt.Printf("Iteration %d ", currentIteration+1)
		fmt.Printf("=> Best value: %E\n", avoa.best1.value)
	}
}

func (avoa *Avoa) exploration(randomSelectIndex int, F float64, currentVulture *vulture) {
	randomSelectVulture := avoa.best1
	if randomSelectIndex == 1 {
		randomSelectVulture = avoa.best2
	}
	r := rand.Float64()
	if rand.Float64() < avoa.p1 {
		for i, pos := range currentVulture.positions {
			currentVulture.positions[i] = randomSelectVulture.positions[i] -
				math.Abs(2*r*randomSelectVulture.positions[i]-pos)*F
		}
	} else {
		r2 := rand.Float64()
		r3 := rand.Float64()
		for i, _ := range currentVulture.positions {
			currentVulture.positions[i] = randomSelectVulture.positions[i] - F +
				r2*((avoa.upperBound[i]-avoa.lowerBound[i])*r3+avoa.lowerBound[i])
		}
	}
}

func (avoa *Avoa) exploitation() {

}
