package main

import (
	"advent_of_code/parser"
	"advent_of_code/utils"
	"fmt"
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

var total = 0
var totals = map[string]int{}
var help = map[string]int{}
var totes = [][]int{}

func walk(e *elFS, path string) {

	var p string
	if e.name == "/" {
		p = path + e.name
	} else {
		p = path + e.name + "/"

	}
	totals[p] = e.size
	// we have reached the end of our walk
	if len(e.next) == 0 {
		// x =  [ asdjf lol ]
		endPath := strings.Split(p, "/")

		for {
			// find the previous key
			key := strings.Join(endPath[0:len(endPath)-1], "/") + "/"

			if key == "/" {
				break
			}
			endPath = endPath[0 : len(endPath)-1]
		}

	}

	for _, c := range e.next {
		walk(c, p)
	}
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
	walk(x, "")
	fmt.Println(len(totals))
	for k, v := range totals {
		if v < 100000 {
			fmt.Println(k, v)

		}
	}
	answer := 0

	for i := 0; i < len(totes); i++ {
		for j := 0; j < len(totes[i])-1; j++ {
			totes[i][j+1] = totes[i][j] + totes[i][j+1]
		}
	}
	for _, x := range totes {
		fmt.Println(x)
		for _, y := range x {
			if y <= 100000 {
				answer += y
			}
		}
	}
	return 0

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
