package main

import (
	"advent_of_code/utils"
	"fmt"
	"strings"
)

func main() {
	input := testInput()
	// input := parser.ReadInputFile(7)
	p1 := part1(input)
	fmt.Println(p1)
}

type elFS struct {
	name string
	next map[string]*elFS
	prev *elFS
	size int
}

var total = 0
var totals = map[string]int{}

func walk(e *elFS) {

	fmt.Println(e.size)
	if len(e.next) == 0 {
		return
	}

	for _, c := range e.next {
		fmt.Printf("walking %s -> %s \n", e.name, c.name)
		walk(c)
	}
}

func part1(input string) int {

	commands := strings.Split(input, "\n")

	dir := &elFS{name: "/", prev: nil, next: map[string]*elFS{}}

	for _, command := range commands[1:] {
		cmd := strings.Split(command, " ")
		// cd into a dir
		if cmd[0] == "$" && cmd[1] == "cd" && cmd[2] != ".." {
			dir.next[cmd[2]] = &elFS{name: cmd[2], prev: dir, next: map[string]*elFS{}}
			dir = dir.next[cmd[2]]
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

	fmt.Println(x)
	fmt.Println(totals)
	return 0

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
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
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
