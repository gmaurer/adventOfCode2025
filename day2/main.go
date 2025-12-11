package main

// this could be better if i didnt use the struct and just immediately
// continued thru the ranges after finding them but i wanted to practice structs.
import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type idRange struct {
	start int
	end   int
}

func main() {
	i, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	total := 0
	allIds := []idRange{}
	stringSlice := strings.Split(string(i), ",")
	for _, val := range stringSlice {
		val = strings.TrimSuffix(val, "\n")
		parts := strings.Split(val, "-")

		start, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		allIds = append(allIds, idRange{
			start: start,
			end:   end,
		})
	}

	for _, idr := range allIds {

		for i := idr.start; i <= idr.end; i++ {
			if len(strconv.Itoa(i))%2 == 0 && strconv.Itoa(i)[:len(strconv.Itoa(i))/2] == strconv.Itoa(i)[len(strconv.Itoa(i))/2:] {
				total += i
			}
		}
	}
	fmt.Println("Total:", total)
}
