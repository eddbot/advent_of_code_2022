package main

import (
	"advent_of_code/parser"
	"advent_of_code/utils"
	"fmt"
	"strings"
)

func main() {
	//	input := testInput()
	input := parser.ReadInputFile(9)
	p1 := part1(input)
	fmt.Println(p1)
}

func part1(input string) int {
	commands := strings.Split(input, "\n")

	// x, y
	headPos := []int{0, 0}
	tailPos := []int{0, 0}
	visited := map[string]bool{}

	posses := [][]int{}

	for _, cmd := range commands {
		c := strings.Split(cmd, " ")
		dir, num := c[0], utils.Conv(c[1])

		for i := 0; i < num; i++ {
			switch dir {
			// if right/left update headpos[0]
			case "R":
				headPos[0]++
			case "L":
				headPos[0]--
			// if up/down update headpos[1]
			case "U":
				headPos[1]++
			case "D":
				headPos[1]--
			}
			tailPos = calculateTail(headPos, tailPos)
			posses = append(posses, tailPos)
			if len(posses) > 9 {
				posses = posses[1:]
			}
			//fmt.Printf("HEAD: %v, TAIL: %v\n", headPos, tailPos)
			//dg := grid()
			//
			//dg[headPos[1]][headPos[0]] = "H"
			//dg[tailPos[1]][tailPos[0]] = "T"
			//
			//for i := len(dg) - 1; i >= 0; i-- {
			//	fmt.Println(dg[i])
			//
			//}
			//fmt.Println("\r\r")
			//time.Sleep(1 * time.Second)
			visited[fmt.Sprintf("x%dy%d", tailPos[0], tailPos[1])] = true
		}

	}
	fmt.Println(len(visited))
	return 0
}

func grid() [][]string {

	g := [][]string{}

	for i := 0; i < 6; i++ {
		gerp := []string{}
		for j := 0; j < 6; j++ {
			gerp = append(gerp, ".")
		}
		g = append(g, gerp)
	}
	return g
}

func calculateTail(headPos, tailPos []int) []int {
	hx, hy := headPos[0], headPos[1]
	tx, ty := tailPos[0], tailPos[1]

	// If we are touching there is no need to move
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			xx := tx + i
			xy := ty + j
			if xx == hx && xy == hy {
				//fmt.Println("touching")
				return tailPos
			}
		}
	}
	// if we get here, need to move
	for i := -2; i <= 2; i++ {
		for j := -2; j <= 2; j++ {
			xx := tx + i
			xy := ty + j
			if xx == hx && xy == hy {
				if (i == -1 && j == -2) || (i == -2 && j == -1) {
					return []int{tailPos[0] - 1, tailPos[1] - 1}
				}
				if (i == 1 && j == -2) || (i == 2 && j == -1) {
					return []int{tailPos[0] + 1, tailPos[1] - 1}
				}
				if (i == 2 && j == 1) || (i == 1 && j == 2) {
					return []int{tailPos[0] + 1, tailPos[1] + 1}
				}
				if (i == -1 && j == 2) || (i == -2 && j == 1) {
					return []int{tailPos[0] - 1, tailPos[1] + 1}
				}
				if j == -2 {
					return []int{tailPos[0], tailPos[1] - 1}
				}
				if j == 2 {
					return []int{tailPos[0], tailPos[1] + 1}
				}
				if i == -2 {
					return []int{tailPos[0] - 1, tailPos[1]}
				}
				if i == 2 {
					return []int{tailPos[0] + 1, tailPos[1]}
				}
			}
		}
	}
	// should never get here
	panic("we lost the head")
}

func testInput() string {
	return `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
}
