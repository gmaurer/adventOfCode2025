package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	startTime := time.Now()
	totalValue := 0

	inputLines := strings.Split(string(input), "\n")
	testSlice := [][]string{}

	for i := range inputLines {
		inputLines[i] = standardizeSpaces(inputLines[i])
		pieces := strings.Split(inputLines[i], " ")
		testSlice = append(testSlice, pieces)
	}
	totalLength := len(strings.Split(inputLines[0], " "))

	for i := range totalLength {
		values := []int{}
		sign := testSlice[len(testSlice)-1][i]

		for j := 0; j <= len(testSlice)-2; j++ {
			if testSlice[j][i] == "" {
				continue
			}
			piece, err := strconv.Atoi(testSlice[j][i])
			if err != nil {
				panic(err)
			}
			values = append(values, piece)
		}

		if sign == "+" {
			totalValue += sum(values)
		} else if sign == "*" {
			totalValue += multiply(values)

		}

	}
	fmt.Println("Total Value: ", totalValue)
	elapsed := time.Since(startTime)
	fmt.Println("Elapsed Time:", elapsed)
}

func sum(values []int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func multiply(values []int) int {
	total := 1
	for _, v := range values {
		total *= v
	}
	return total
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
