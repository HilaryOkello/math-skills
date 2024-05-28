package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"math-skills/formulas"
)

func main() {
	// Check command for correctness
	if len(os.Args) != 2 {
		log.Fatal("Incorrect command.\nExpects: \"go run main.go <filename>\"")
	}
	fileName := os.Args[1]
	if fileName[len(fileName)-4:] != ".txt" {
		log.Fatal("Wrong file extension. Expects a .txt file")
	}

	// Opening file, defer close, and check if it's empty
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

	// create a scanner, read each line, and append nums to []float64
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
	// call functions from the formulas package to calculate our statistics, then print
	average := formulas.CalcAverage(dataInts)
	variance := formulas.CalcVariance(dataInts, average)
	median := formulas.CalcMedian(dataInts)
	standardDeviation := formulas.CalcStdDev(variance)

	fmt.Printf("Average: %.f\nMedian: %.f\nVariance: %.f\nStandard Deviation: %.f\n",
	average, median, variance, standardDeviation)
}
