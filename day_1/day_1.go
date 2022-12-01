package main

import (
	"advent_of_code/parser"
	"fmt"
	"strconv"
	"strings"
)

var solutions = make(map[int]int)

func main() {
	optimizedSolution()
	fmt.Println("part 1", solutions[1])
	fmt.Println("part 2", solutions[2])
}

func optimizedSolution() {
	input := parser.ReadInputFile(1)
	elves := strings.Split(input, "\n\n")
	lo, mid, hi := 0, 0, 0

	for _, elf := range elves {
		cals := strings.Split(elf, "\n")
		totalCals := 0
		for _, cal := range cals {
			c, err := strconv.Atoi(cal)
			if err != nil {
				panic(err)
			}
			totalCals += c
		}

		// for part 1
		if totalCals > solutions[1] {
			solutions[1] = totalCals
		}

		// for part 2
		switch {
		case totalCals > hi:
			lo = mid
			mid = hi
			hi = totalCals
		case totalCals > mid:
			lo = mid
			mid = totalCals
		case totalCals > lo:
			lo = totalCals
		}
	}
	solutions[2] = lo + mid + hi
}
