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

func solution(input string)(int,int){
	total1, total2 := 0,0
	noOverlaps := 0
	assignments := strings.Split(input, "\n")
	for _, assignment := range assignments {
		gnomes := strings.Split(assignment, ",")
		gnomeA, gnomeB := gnomes[0], gnomes[1]
		gnomeAMin, gnomeAMax := minMax(gnomeA)
		gnomeBMin, gnomeBMax := minMax(gnomeB)

		total2++
		if cmp1(gnomeAMin, gnomeBMin, gnomeAMax, gnomeBMax) {
			total1++
		}
		if cmp2(gnomeAMin, gnomeBMin, gnomeAMax, gnomeBMax) {
			noOverlaps++
		}
	}
	return total1, total2 - noOverlaps
}

func cmp1(minA, minB, maxA, maxB int)bool {
	return minA <= minB && maxA >= maxB ||
	minB <= minA && maxB >= maxA
}

func cmp2(minA, minB, maxA, maxB int)bool {
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
