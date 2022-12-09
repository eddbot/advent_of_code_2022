package main

import (
	"advent_of_code/parser"
	"advent_of_code/utils"
	"fmt"
	"math"
	"strings"
)

func main() {
	// input := testInput()
	input := parser.ReadInputFile(7)

	p1 := part1(input)
	fmt.Println(p1)
}

type elFS struct {
	name string
	next map[string]*elFS
	prev *elFS
	size int
}

var totals = []int{}

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

func part1(input string) int {

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

	x := rewind(dir)
	walk(x)

	// answer := 0

	// part 1
	// for _, t := range totals {
	// 	if t < 100000 {
	// 		answer += t
	// 	}
	// }
	// fmt.Println(answer)

	totalSpace := 70000000

	unusedSpace := (totalSpace - x.size)
	requiredSpace := 30000000 - unusedSpace

	smallest := math.MaxInt

	for _, t := range totals {
		if t > requiredSpace {
			if t < smallest {
				smallest = t
			}
		}
	}
	return smallest

}

func puts[T comparable](in T) {
	fmt.Println(in)
}

func rewind(e *elFS) *elFS {
	if e.prev == nil {
		return e
	}
	return rewind(e.prev)
}
func testInput() string {
	return `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd aaaa
$ cd lol
$ dir fucku
$ cd ..
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd eee
$ cd empty
$ ls
dir zz
$ cd ..
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`
}
