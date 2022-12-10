package main

import (
	"advent_of_code/parser"
	"advent_of_code/utils"
	"fmt"
	"strings"
)

// 1816
// 383520
func main() {
	input := parser.ReadInputFile(8)
	forest := createForest(input)
	p1 := part1(forest)
	p2 := part2(forest)
	fmt.Println(p1)
	fmt.Println(p2)
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

func part2(trees [][]int) int {

	locations := [][]int{}

	for i := 1; i < len(trees)-1; i++ {
		for j := 1; j < len(trees[i])-1; j++ {
			// lets not even consider the edges, they will be 0 so nein
			height := trees[i][j]
			scores := []int{}

			scenicScore := 0

			// look up
			ups := i
			for {
				ups--
				// we hit an edge
				if ups < 0 {
					break
				}
				scenicScore++
				// tree too high
				if trees[ups][j] >= height {
					break
				}

			}
			scores = append(scores, scenicScore)
			scenicScore = 0

			downs := i
			for {
				downs++

				// we hit an edge
				if downs > len(trees[i])-1 {
					break
				}
				scenicScore++
				if trees[downs][j] >= height {
					break
				}
			}
			scores = append(scores, scenicScore)

			// look left
			scenicScore = 0

			lefts := j
			for {
				lefts--
				// we hit an edge
				if lefts < 0 {
					break
				}
				scenicScore++
				if trees[i][lefts] >= height {
					break
				}

			}
			scores = append(scores, scenicScore)

			scenicScore = 0

			rights := j
			for {
				rights++
				// we hit an edge
				if rights > len(trees)-1 {
					break
				}
				scenicScore++
				if trees[i][rights] >= height {
					break
				}
			}
			scores = append(scores, scenicScore)

			locations = append(locations, scores)
		}

	}

	answer := 0

	for _, location := range locations {
		acc := accumulate(location)
		if acc > answer {
			answer = acc
		}
	}
	return answer
}

func accumulate(scores []int) int {
	score := scores[0]

	for _, tree := range scores[1:] {
		score *= tree
	}
	return score
}
