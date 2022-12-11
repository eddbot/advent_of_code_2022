package main

import (
	"advent_of_code/parser"
	"advent_of_code/utils"
	"fmt"
	"strings"
)

func main() {

	//input := testInput()
	// input := bigtestInput()
	input := parser.ReadInputFile(10)
	p1 := part1(input)
	fmt.Println(p1)
}

func part1(input string) int {
	register := 1
	totalCycles := 0
	//sigger := 0
	j := 0
	sigStrengths := []int{}

	screen := crtScreen()

	instructions := strings.Split(input, "\n")

	for _, instruction := range instructions {
		op := strings.Split(instruction, " ")
		exe := false
		var clockCycles int

		if len(op) == 1 {
			// noop
			clockCycles = 1
		} else {
			exe = true
			clockCycles = 2
		}

		for i := clockCycles; i > 0; i-- {

			// part 1
			//if totalCycles == 20+sigger {
			//	sigger += 40
			//	sigStrengths = append(sigStrengths, register*totalCycles)
			//}

			// part 2
			if register == (totalCycles-(j*40)) || register-1 == (totalCycles-(j*40)) || register+1 == (totalCycles-(j*40)) {
				screen[j][(totalCycles - (j * 40))] = "#"
			}
			totalCycles++
			fmt.Println(j, totalCycles, register)
			if totalCycles%40 == 0 {
				j++
			}
		}
		if exe {
			register += utils.Conv(op[1])
		}
	}

	for _, row := range screen {
		fmt.Println(row)
	}
	return accumulate(sigStrengths)
}

func accumulate(sigs []int) int {
	total := 0

	for _, v := range sigs {
		total += v
	}
	return total
}

func crtScreen() [][]string {
	crt := [][]string{}

	for i := 0; i < 6; i++ {
		c := []string{}
		for i := 0; i < 40; i++ {
			c = append(c, ".")
		}
		crt = append(crt, c)
	}

	return crt
}
func testInput() string {
	return `noop
addx 3
addx -5`
}

func bigtestInput() string {
	return `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`
}
