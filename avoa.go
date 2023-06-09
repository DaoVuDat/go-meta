package main

import (
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
	for i := range avoa.population {
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
		for i := range avoa.population {
			vult := &avoa.population[i]
			F := P1 * (2.0*rand.Float64() - 1)

			// random select vulture
			randomSelectVultureIndex := RWS([]float64{avoa.alpha, avoa.betha})

			if math.Abs(F) >= 1 {
				// Exploration -> update position
				avoa.exploration(randomSelectVultureIndex, F, vult)
			} else {
				// Exploitation -> update position
				avoa.exploitation(randomSelectVultureIndex, F, vult)
			}

			vult.value = avoa.objectiveFunction(vult.positions)
		}

		// Find top 2 vultures
		avoa.findBest()

		// Print information
		//fmt.Printf("Iteration %d ", currentIteration+1)
		//fmt.Printf("=> Best value: %E\n", avoa.best1.value)
	}
}

func boundaryChecking(newPos float64, upperBound float64, lowerBound float64) float64 {
	if newPos > upperBound {
		newPos = upperBound
	} else if newPos < lowerBound {
		newPos = lowerBound
	}

	return newPos
}

func (avoa *Avoa) exploration(randomSelectIndex int, F float64, currentVulture *vulture) {
	randomSelectVulture := avoa.best1
	if randomSelectIndex == 1 {
		randomSelectVulture = avoa.best2
	}
	r := rand.Float64()
	if rand.Float64() < avoa.p1 {
		for i, pos := range currentVulture.positions {
			newPos := randomSelectVulture.positions[i] -
				math.Abs(2*r*randomSelectVulture.positions[i]-pos)*F

			currentVulture.positions[i] = boundaryChecking(newPos, avoa.upperBound[i], avoa.lowerBound[i])
		}
	} else {
		r2 := rand.Float64()
		r3 := rand.Float64()
		for i := range currentVulture.positions {
			newPos := randomSelectVulture.positions[i] - F +
				r2*((avoa.upperBound[i]-avoa.lowerBound[i])*r3+avoa.lowerBound[i])

			currentVulture.positions[i] = boundaryChecking(newPos, avoa.upperBound[i], avoa.lowerBound[i])
		}
	}
}

func (avoa *Avoa) exploitation(randomSelectIndex int, F float64, currentVulture *vulture) {
	randomSelectVulture := avoa.best1
	if randomSelectIndex == 1 {
		randomSelectVulture = avoa.best2
	}
	if math.Abs(F) < 0.5 {
		if rand.Float64() < avoa.p2 {
			for i, pos := range currentVulture.positions {
				A := avoa.best1.positions[i] - ((avoa.best1.positions[i]*pos)/(avoa.best1.positions[i]-pos*pos))*F
				B := avoa.best2.positions[i] - ((avoa.best2.positions[i]*pos)/(avoa.best2.positions[i]-pos*pos))*F
				curPos := (A + B) / 2
				currentVulture.positions[i] = boundaryChecking(curPos, avoa.upperBound[i], avoa.lowerBound[i])
			}
		} else {
			levy := LevyFlight(avoa.dimension)
			for i, pos := range currentVulture.positions {
				newPos := randomSelectVulture.positions[i] -
					math.Abs(randomSelectVulture.positions[i]-pos)*F*levy[i]

				currentVulture.positions[i] = boundaryChecking(newPos, avoa.upperBound[i], avoa.lowerBound[i])
			}
		}
	} else {
		if rand.Float64() < avoa.p3 {
			r1 := rand.Float64()
			r2 := rand.Float64()
			for i, pos := range currentVulture.positions {
				newPos := (math.Abs(2.0*r1*randomSelectVulture.positions[i])-pos)*
					(F+r2) - (randomSelectVulture.positions[i] - pos)
				currentVulture.positions[i] = boundaryChecking(newPos, avoa.upperBound[i], avoa.lowerBound[i])
			}
		} else {
			for i, pos := range currentVulture.positions {
				s1 := randomSelectVulture.positions[i] * (rand.Float64() * pos / (2.0 * math.Pi) * math.Cos(pos))
				s2 := randomSelectVulture.positions[i] * (rand.Float64() * pos / (2.0 * math.Pi) * math.Sin(pos))
				newPos := randomSelectVulture.positions[i] - (s1 + s2)
				currentVulture.positions[i] = boundaryChecking(newPos, avoa.upperBound[i], avoa.lowerBound[i])
			}
		}
	}
}
