package main

import (
	"advent_of_code/parser"
	"fmt"
)

func main() {
	input := parser.ReadInputFile(5)

	stacks, instructions := puzzleParser(input)
	fmt.Println(part1(stacks, instructions))

	stacks, instructions = puzzleParser(input)
	fmt.Println(part2(stacks, instructions))

}

func part1(stacks map[int][]string, instructions []instruction) string {
	for _, cmd := range instructions {
		for i := 0; i < cmd.move; i++ {
			stacks[cmd.to] = insert(stacks[cmd.to], stacks[cmd.from][i])
		}
		stacks[cmd.from] = stacks[cmd.from][cmd.move:]
	}
	return answerParser(stacks)
}

func part2(stacks map[int][]string, instructions []instruction) string {
	for _, cmd := range instructions {
		stacks[cmd.to] = insertMulti(stacks[cmd.to], stacks[cmd.from][0:cmd.move])
		stacks[cmd.from] = stacks[cmd.from][cmd.move:]
	}
	return answerParser(stacks)
}
