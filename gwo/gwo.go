package gwo

import "gonum.org/v1/gonum/stat"

func Test(x []float64) float64 {
	weights := []float64{1, 1, 1, 1, 1}
	return stat.Mean(x, weights)
}
