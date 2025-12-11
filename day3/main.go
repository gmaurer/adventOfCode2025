package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	totalJoltage := 0
	i, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	batteries := []string{}
	stringSlice := strings.Split(string(i), "\n")

	batteries = append(batteries, stringSlice...)

	for _, val := range batteries {
		var tempFirstIndex, tempFirst int
		var _, tempSecond int

		tempFirstIndex, tempFirst = helperFunction(val[:len(val)-1])
		_, tempSecond = helperFunction(val[tempFirstIndex+1:])
		finalInt, err := strconv.Atoi(strconv.Itoa(tempFirst) + strconv.Itoa(tempSecond))
		if err != nil {
			panic(err)
		}
		totalJoltage += finalInt

	}
	fmt.Println("Total Joltage:", totalJoltage)
}

func helperFunction(strArr string) (int, int) {
	var tempVal, tempIndex int
	tempVal = 0
	tempIndex = 0
	for i, char := range strArr {
		charInt, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		if charInt == 9 {
			return i, charInt
		}
		if charInt > tempVal {
			tempVal = charInt
			tempIndex = i
		}
	}
	return tempIndex, tempVal
}
