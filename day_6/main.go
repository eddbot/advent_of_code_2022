package main

import (
	"advent_of_code/parser"
	"fmt"
)

func main() {
	input := parser.ReadInputFile(6)
	fmt.Println(solver(input, 4))
	fmt.Println(solver(input, 14))
}

func solver(input string, uniqToFind int) int {
	for i := 0; i < len(input)-(uniqToFind-1); i++ {
		dups := map[byte]bool{}
		for j := 0; j < uniqToFind; j++ {
			dups[input[i+j]] = true
		}
		if len(dups) == uniqToFind {
			return i + uniqToFind
		}
	}
	return 0
}
