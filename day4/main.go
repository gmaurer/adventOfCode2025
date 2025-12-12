package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	x := strings.Split(string(input), "\n")

	grid := [][]string{}
	for _, line := range x {
		row := strings.Split(line, "")
		grid = append(grid, row)
	}
	accessableToiletPaper := 0
	for dx := range grid {
		for dy := range grid[dx] {
			numToiletPaper := 0
			if grid[dx][dy] == "." {
				continue
			}
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if i == 0 && j == 0 {
						continue
					}
					if dx+i < 0 || dx+i >= len(grid) || dy+j < 0 || dy+j >= len(grid[dx]) {
						continue
					}
					if grid[dx+i][dy+j] == "@" {
						numToiletPaper++
					}
				}
			}
			if numToiletPaper < 4 {
				accessableToiletPaper++
			}
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Accessable Toilet Paper:", accessableToiletPaper)
	fmt.Println("Elapsed Time:", elapsed)
}
