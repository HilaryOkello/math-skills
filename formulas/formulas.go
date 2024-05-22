package formulas

import (
	"math"
	"sort"
)
//Receives our data []float64 and returns the average
func CalcAverage(data []float64) float64 {
	var sum float64
	for _, num := range data {
		sum += num
	}
	return sum / float64(len(data))
}
//Receives our data []float64 and returns the median
func CalcMedian(data []float64) float64 {
	//sorts the []float in ascending order
	sort.SliceStable(data, func(i, j int) bool { return data[i] < data[j] })
	if len(data)%2 != 0 {
		return data[len(data)/2]
	}
	sumOfMiddle := data[len(data)/2] + data[(len(data)/2)-1]
	return sumOfMiddle / 2
}
//Receives our data []float64 and the average, calculates the variance and returns it
func CalcVariance(data []float64, mean float64) float64 {
	var sumOfSquares float64
	for _, num := range data {
		diff := num - mean
		sumOfSquares += math.Pow(diff, 2)
	}
	return sumOfSquares / float64(len(data))
}
//Squareroots the variance and returns the Standard Deviation
func CalcStdDev(variance float64) float64 {
	return math.Sqrt(variance)
}
