package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type idRange struct {
	start int
	end   int
}

func main() {
	startTime := time.Now()
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	idToCheck := []int{}
	rangesOf := []idRange{}
	fresh := 0

	inputLines := strings.Split(string(input), "\n\n")

	rangesInput := strings.Split(inputLines[0], "\n")
	ids := strings.Split(inputLines[1], "\n")
	for _, idStr := range ids {
		if idStr == "" {
			continue
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}
		idToCheck = append(idToCheck, id)
	}
	for _, line := range rangesInput {
		parts := strings.Split(line, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		rangesOf = append(rangesOf, idRange{
			start: start,
			end:   end,
		})

	}

	for _, i := range idToCheck {
		if helperFunction(i, rangesOf) {
			fresh += 1
		}
	}
	elapsed := time.Since(startTime)
	fmt.Println("Fresh Ingreds:", fresh)
	fmt.Println("Elapsed Time:", elapsed)
}

func helperFunction(id int, ranges []idRange) bool {
	for _, r := range ranges {
		start := r.start
		end := r.end
		if id >= start && id <= end {
			return true
		}
	}
	return false
}
