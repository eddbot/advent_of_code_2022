package main

import (
	"advent_of_code/parser"
	"advent_of_code/utils"
	"fmt"
	"math"
	"strings"
)

func main() {
	input := parser.ReadInputFile(7)
	fs := createFS(input)
	p1 := part1(fs)
	p2 := part2(fs)
	fmt.Printf("Part 1: %d, Part 2: %d\n", p1, p2)
}

type elFS struct {
	name string
	next map[string]*elFS
	prev *elFS
	size int
}

var totals = []int{}

func rewind(e *elFS) *elFS {
	if e.prev == nil {
		return e
	}
	return rewind(e.prev)
}

func walk(e *elFS) int {
	// we have reached the end of our walk
	if len(e.next) == 0 {
		totals = append(totals, e.size)
		return e.size
	}
	for _, c := range e.next {
		e.size += walk(c)
	}
	totals = append(totals, e.size)
	return e.size
}

func createFS(input string) *elFS {
	commands := strings.Split(input, "\n")
	dir := &elFS{name: "/", prev: nil, next: map[string]*elFS{}}

	for _, command := range commands[1:] {
		cmd := strings.Split(command, " ")
		// cd into a dir
		if cmd[0] == "$" && cmd[1] == "cd" && cmd[2] != ".." {
			// cd into existing dir
			if d, ok := dir.next[cmd[2]]; ok {
				dir = d
			} else {
				// create a new dir
				dir.next[cmd[2]] = &elFS{name: cmd[2], prev: dir, next: map[string]*elFS{}}
				dir = dir.next[cmd[2]]
			}
		}
		// cd out of a dir
		if cmd[0] == "$" && cmd[1] == "cd" && cmd[2] == ".." {
			dir = dir.prev
		}
		// do the sizes
		if cmd[0] != "$" && cmd[0] != "dir" {
			size := utils.Conv(cmd[0])
			dir.size += size
		}
	}
	return dir
}

func part1(dir *elFS) int {
	x := rewind(dir)
	walk(x)
	answer := 0
	for _, t := range totals {
		if t < 100000 {
			answer += t
		}
	}
	return answer
}

func part2(dir *elFS) int {
	x := rewind(dir)
	requiredSpace := 30000000 - (70000000 - x.size)
	answer := math.MaxInt
	for _, t := range totals {
		if t > requiredSpace {
			if t < answer {
				answer = t
			}
		}
	}
	return answer
}
