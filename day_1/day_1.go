package main

import (
	"advent_of_code/parser"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	cals := extractCalories()
	fmt.Println("part 1", part1(cals))
	fmt.Println("part 2", part2(cals))
}

func part1(cals []int) int {
	return cals[len(cals)-1]
}

func part2(cals []int) int {
	i := len(cals) - 1
	return cals[i] + cals[i-1] + cals[i-2]
}

func extractCalories() []int {
	input := parser.ReadInputFile(1)
	elves := strings.Split(input, "\n\n")
	calories := []int{}

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
		calories = append(calories, totalCals)
	}

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] < calories[j]
	})
	return calories
}
