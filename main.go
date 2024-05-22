package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"math-skills/formulas"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Incorrect command.\nExpects: \"go run main.go <filename>\"")
	}
	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening %s:\n%s", fileName, err)
	}
	defer file.Close()

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fileSize := fileInfo.Size()

	if fileSize == 0 {
		log.Fatal(fileName, " is empty.")
	}

	scanner := bufio.NewScanner(file)
	var dataInts []float64
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Fatalf("%s may contain non number data: %s", fileName, err)
		}
		dataInts = append(dataInts, num)
	}

	average := math.Round(formulas.CalcAverage(dataInts))
	median := math.Round(formulas.CalcMedian(dataInts))
	variance := math.Round(formulas.CalcVariance(dataInts, average))
	standardDeviation := math.Round(formulas.CalcStdDev(variance))

	fmt.Printf("Average: %.0f\nMedian: %.0f\nVariance: %.0f\nStandard Deviation: %.0f\n",
		average, median, variance, standardDeviation)
}
