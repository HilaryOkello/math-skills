package main

import (
	"fmt"
	"math/big"
	"os"
	"sort"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	dataString := strings.Split(string(content), "\n")
	var dataInts []big.Int
	for _, str := range dataString {
		var bigNum big.Int
		_, success := bigNum.SetString(str, 10)
		if !success {
			fmt.Println("unable to convert string to big int")
		}
		dataInts = append(dataInts, bigNum)
	}
	mean := calculateAverage(dataInts)
	median := calculateMedian(dataInts)
	fmt.Println(mean, median)
}

func calculateAverage(data []big.Int) *big.Int {
	var sum big.Int
	for _, num := range data {
		sum.Add(&sum, &num)
	}
	count := big.NewInt(int64(len(data)))
	var mean big.Int
	mean.Div(&sum, count)
	return &mean
}

func calculateMedian(data []big.Int) *big.Int {
	sort.SliceStable(data, func(i, j int) bool { return i < j })
	if len(data)%2 != 0 {
		return &data[len(data)/2]
	}
	var sum big.Int
	sum.Add(&data[len(data)/2], &data[(len(data)/2)-1])
	var median big.Int
	median.Div(&sum, big.NewInt(2))
	return &median
}
