package main

import (
	"advent_of_code/parser"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := parser.ReadInputFile(4)
	part1, part2 := solution(input)
	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}

func solution(input string) (int, int) {
	part1, part2, noOverlaps := 0, 0, 0
	assignments := strings.Split(input, "\n")
	for _, assignment := range assignments {
		gnomes := strings.Split(assignment, ",")
		gnomeAMin, gnomeAMax := minMax(gnomes[0])
		gnomeBMin, gnomeBMax := minMax(gnomes[1])

		if fullOverlap(gnomeAMin, gnomeBMin, gnomeAMax, gnomeBMax) {
			part1++
		}
		if noOverlap(gnomeAMin, gnomeBMin, gnomeAMax, gnomeBMax) {
			noOverlaps++
		}
		part2++
	}
	return part1, (part2 - noOverlaps)
}

func fullOverlap(minA, minB, maxA, maxB int) bool {
	return minA <= minB && maxA >= maxB ||
		minB <= minA && maxB >= maxA
}

func noOverlap(minA, minB, maxA, maxB int) bool {
	return maxA < minB || maxB < minA
}

func minMax(gnome string) (int, int) {
	g := strings.Split(gnome, "-")

	min, err := strconv.Atoi(g[0])
	if err != nil {
		panic(err)
	}
	max, err := strconv.Atoi(g[1])
	if err != nil {
		panic(err)
	}
	return min, max
}
