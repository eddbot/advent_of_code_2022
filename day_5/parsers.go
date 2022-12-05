package main

import (
	"strings"
)

type instruction struct {
	move int
	from int
	to   int
}

type stacks map[int][]string

func puzzleParser(input string) (stacks, []instruction) {
	in := strings.Split(input, "\n\n")
	return extractStacks(in[0]), extractInstructions(in[1])
}

func extractStacks(input string) stacks {
	rows := strings.Split(input, "\n")
	out := stacks{}
	for i, row := range rows {
		col := 1
		for j := 1; j < len(row)-1; j += 4 {
			cc := string(row[j])
			if cc == " " || i == len(rows)-1 {
				col++
				continue
			}
			out[col] = append(out[col], string(row[j]))
			col++
		}
	}
	return out
}

func extractInstructions(input string) []instruction {
	instructions := []instruction{}
	rows := strings.Split(input, "\n")

	for _, row := range rows {
		r := strings.Split(row, " ")
		instructions = append(instructions, instruction{
			move: conv(r[1]),
			from: conv(r[3]),
			to:   conv(r[5]),
		})
	}
	return instructions
}
