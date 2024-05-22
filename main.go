package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"math-skills/formulas"
)

func main() {
	fileName := os.Args[1]
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	dataString := strings.Split(string(content[:len(content)-1]), "\n")
	var dataInts []float64
	for _, str := range dataString {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err)
			return
		}
		dataInts = append(dataInts, float64(num))
	}

	average := math.Round(formulas.CalcAverage(dataInts))
	median := math.Round(formulas.CalcMedian(dataInts))
	variance := math.Round(formulas.CalcVariance(dataInts, average))
	standardDeviation := math.Round(formulas.CalcStdDev(variance))

	fmt.Printf("Average: %.0f\nMedian: %.0f\nVariance: %.0f\nStandard Deviation: %.0f\n",
		average, median, variance, standardDeviation)
}
