package main

import (
	"advent_of_code/parser"
	"advent_of_code/utils"
	"fmt"
	"strings"
)

func main() {
	//input := testInput()
	input := parser.ReadInputFile(8)
	forest := createForest(input)
	p1 := part1(forest)
	fmt.Println(p1)
}

func createForest(input string) [][]int {
	trees := [][]int{}

	rows := strings.Split(input, "\n")
	for _, row := range rows {
		t := strings.Split(row, "")
		tt := []int{}

		for _, tr := range t {
			tt = append(tt, utils.Conv(tr))
		}
		trees = append(trees, tt)
	}
	return trees
}

func part1(trees [][]int) int {
	visibleTrees := 0
	fuu := map[string]int{}

	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			// easy part lol - outside trees
			if i == 0 || i == len(trees)-1 || j == 0 || j == len(trees[i])-1 {
				visibleTrees++
				continue
			}
			// check top
			height := trees[i][j]
			startPos := 0
			currentPos := i
			for {
				// we made it to a tree!
				if startPos == currentPos {
					fu := fmt.Sprintf("[%d, %d]", startPos, j)
					fuu[fu] = 1
					break

				}
				// tree_too_high.jpg
				if trees[startPos][j] >= height {
					break
				}
				// move to the next tree
				startPos++
			}

			// check bottom
			height = trees[i][j]
			startPos = len(trees[i]) - 1
			currentPos = i
			for {
				// we made it to a tree!
				if startPos == currentPos {
					fu := fmt.Sprintf("[%d, %d]", startPos, j)
					fuu[fu] = 1
					break
				}
				// tree_too_high.jpg
				if trees[startPos][j] >= height {
					break
				}
				// move to the next tree
				startPos--
			}

			// check left
			height = trees[i][j]
			startPos = len(trees[i]) - 1
			currentPos = j
			for {
				// we made it to a tree!
				if startPos == currentPos {
					fu := fmt.Sprintf("[%d, %d]", i, startPos)
					fuu[fu] = 1
					break
				}
				// tree_too_high.jpg
				if trees[i][startPos] >= height {
					break
				}
				// move to the next tree
				startPos--
			}
			height = trees[i][j]
			startPos = 0
			currentPos = j
			for {
				// we made it to a tree!
				if startPos == currentPos {
					fu := fmt.Sprintf("[%d, %d]", i, startPos)
					fuu[fu] = 1
					break
				}
				// tree_too_high.jpg
				if trees[i][startPos] >= height {
					break
				}
				// move to the next tree
				startPos++
			}
		}
	}

	for _, v := range fuu {
		visibleTrees += v
	}
	return visibleTrees
}

func part2(trees [][]int) int {}

func testInput() string {
	return `30373
25512
65332
33549
35390`
}
