package main

import (
	"advent_of_code/parser"
	"fmt"
	"strings"
)

func main() {

	input := parser.ReadInputFile(2)
	solutions := map[int]int{1: 0, 2: 0}

	part1 := map[string]map[string]int{
		"A": {"X": 4, "Y": 8, "Z": 3},
		"B": {"X": 1, "Y": 5, "Z": 9},
		"C": {"X": 7, "Y": 2, "Z": 6},
	}
	part2 := map[string]map[string]int{
		"A": {"X": 3, "Y": 4, "Z": 8},
		"B": {"X": 1, "Y": 5, "Z": 9},
		"C": {"X": 2, "Y": 6, "Z": 7},
	}
	strategyGuide := strings.Split(input, "\n")

	for _, t := range strategyGuide {
		turn := strings.Split(t, " ")
		cpu, me := turn[0], turn[1]
		solutions[1] += part1[cpu][me]
		solutions[2] += part2[cpu][me]
	}

	fmt.Println("part 1:", solutions[1])
	fmt.Println("part 2:", solutions[2])
}
