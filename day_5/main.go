package main

import (
	"advent_of_code/parser"
	"fmt"
)

func main() {
	input := parser.ReadInputFile(5)
	stacks, instructions := puzzleParser(input)

	fmt.Println(part1(stacks.dup(), instructions))
	fmt.Println(part2(stacks.dup(), instructions))
}

func part1(stacks stacks, instructions []instruction) string {
	for _, cmd := range instructions {
		move, to, from := cmd.populate()
		for i := 0; i < move; i++ {
			stacks[to] = stacks.insert(to, from, i)
		}
		stacks[from] = stacks.remain(from, move)
	}
	return answerParser(stacks)
}

func part2(stacks stacks, instructions []instruction) string {
	for _, cmd := range instructions {
		move, to, from := cmd.populate()
		stacks[to] = stacks.insertMultiple(to, from, move)
		stacks[from] = stacks.remain(from, move)
	}
	return answerParser(stacks)
}
