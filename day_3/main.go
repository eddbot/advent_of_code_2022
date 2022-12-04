package main

import (
	"advent_of_code/parser"
	"fmt"
	"strings"
)

func main() {
	input := parser.ReadInputFile(3)
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}

func part1(input string) int {
	rucksacks := strings.Split(input, "\n")
	total := 0
	for _, rucksack := range rucksacks {
		c := len(rucksack)
		items := strings.Split(rucksack, "")
		compartment1, compartment2 := items[0:c/2], items[c/2:]
		ci := findCommonItem(compartment1, compartment2)

		total += letterParser(ci)
	}
	return total
}

func part2(input string) int {
	rucksacks := strings.Split(input, "\n")
	total := 0

	for i := 0; i < len(rucksacks)-2; i += 3 {
		itemCache := map[rune]int{}
		line := uniq(rucksacks[i]) + uniq(rucksacks[i+1]) + uniq(rucksacks[i+2])

		for _, x := range line {
			itemCache[x]++
			if itemCache[x] == 3 {
				total += letterParser(string(x))
			}
		}
	}
	return total
}

func findCommonItem(c1, c2 []string) string {
	common := map[string]bool{}

	for i := 0; i < len(c1); i++ {
		common[c1[i]] = true
	}
	for i := 0; i < len(c2); i++ {
		if ok := common[c2[i]]; ok {
			return c2[i]
		}
	}
	return ""
}

func uniq(input string) string {
	uniq := map[rune]bool{}
	for _, s := range input {
		uniq[s] = true
	}
	uniqString := ""
	for k := range uniq {
		uniqString += string(k)
	}
	return uniqString
}

func letterParser(in string) int {
	cd := []byte(in)[0]
	if cd > 90 {
		return int(cd - 96)
	}
	return int(cd - 38)
}
