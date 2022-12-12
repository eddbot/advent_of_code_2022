package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	input := testInput()

	p1 := part1(input)
	fmt.Println(p1)

}

func walk(x, y, steps int, prev rune, mp [][]rune) {

	time.Sleep(1 * time.Second)
	fmt.Println(x, y)

	// check if ob
	if x < 0 || x > len(mp[0])-1 {
		return
	}
	if y < 0 || y > len(mp[1])-1 {
		return
	}

	// we reached the end
	if mp[x][y] == 69 {
		return
	}

	// too high to climb
	if mp[x][y] > prev && prev != 83 {
		return
	}

	fmt.Println(steps)

	fmt.Println(mp[x][y])

	me := mp[x][y]

	walk(x-1, y, steps+1, me, mp)
	walk(x, y-1, steps+1, me, mp)
	walk(x+1, y, steps+1, me, mp)
	walk(x, y+1, steps+1, me, mp)

}

func createMap(input string) ([][]rune, []int, []int) {

	rows := strings.Split(input, "\n")
	out := [][]rune{}
	var start, end []int

	for i, row := range rows {
		rr := []rune{}

		for j, r := range row {
			if r == 'S' {
				start = []int{i, j}
			}
			if r == 'E' {
				end = []int{i, j}
			}
			rr = append(rr, r)
		}
		out = append(out, rr)

	}
	return out, start, end
}

func part1(input string) int {
	mm, start, _ := createMap(input)

	for _, row := range mm {
		fmt.Println(row)
	}

	walk(start[0], start[1], 0, 83, mm)
	return 0
}

func testInput() string {
	return `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
}
