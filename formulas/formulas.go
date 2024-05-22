package formulas

import (
	"math"
	"sort"
)

func CalcAverage(data []float64) float64 {
	var sum float64
	for _, num := range data {
		sum += num
	}
	n := float64(len(data))
	average := sum / n
	return average
}

func CalcMedian(data []float64) float64 {
	sort.SliceStable(data, func(i, j int) bool { return data[i] < data[j] })
	if len(data)%2 != 0 {
		return data[len(data)/2]
	}
	sumOfMiddle := data[len(data)/2] + data[(len(data)/2)-1]
	median := sumOfMiddle / 2
	return median
}

func CalcVariance(data []float64, mean float64) float64 {
	var sumOfSquares float64
	for _, num := range data {
		diff := num - mean
		sumOfSquares += math.Pow(diff, 2)
	}
	variance := sumOfSquares / float64(len(data))
	return variance
}

func CalcStdDev(variance float64) float64 {
	return math.Sqrt(variance)
}
