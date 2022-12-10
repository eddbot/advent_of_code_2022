package main

import (
	"advent_of_code/parser"
	"advent_of_code/utils"
	"fmt"
	"strings"
)

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
	treeScanner := map[string]int{}

	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			// easy part lol - outside trees
			if i == 0 || i == len(trees)-1 || j == 0 || j == len(trees[i])-1 {
				visibleTrees++
				continue
			}
			height := trees[i][j]

			// check top
			for startPos := 0; startPos <= i; startPos++ {
				if startPos == i {
					t := fmt.Sprintf("[%d, %d]", startPos, j)
					treeScanner[t] = 1
					break
				}
				if trees[startPos][j] >= height {
					break
				}
			}
			// check bottom
			for startPos := len(trees[i]) - 1; startPos >= i; startPos-- {
				if startPos == i {
					t := fmt.Sprintf("[%d, %d]", startPos, j)
					treeScanner[t] = 1
					break
				}
				if trees[startPos][j] >= height {
					break
				}
			}
			// check left
			for startPos := len(trees[i]) - 1; startPos >= j; startPos-- {
				if startPos == j {
					t := fmt.Sprintf("[%d, %d]", i, startPos)
					treeScanner[t] = 1
					break
				}
				if trees[i][startPos] >= height {
					break
				}
			}
			// check right
			for startPos := 0; startPos <= j; startPos++ {
				if startPos == j {
					t := fmt.Sprintf("[%d, %d]", i, startPos)
					treeScanner[t] = 1
					break
				}
				if trees[i][startPos] >= height {
					break
				}
			}
		}
	}
	for _, v := range treeScanner {
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
			for ups := (i - 1); ups >= 0; ups-- {
				scenicScore++
				if trees[ups][j] >= height {
					break
				}
			}
			scores = append(scores, scenicScore)

			// look down
			scenicScore = 0
			for downs := (i + 1); downs < len(trees); downs++ {
				scenicScore++
				if trees[downs][j] >= height {
					break
				}
			}
			scores = append(scores, scenicScore)

			// look left
			scenicScore = 0
			for lefts := (j - 1); lefts >= 0; lefts-- {
				scenicScore++
				if trees[i][lefts] >= height {
					break
				}
			}
			scores = append(scores, scenicScore)

			// look right
			scenicScore = 0
			for rights := (j + 1); rights < len(trees); rights++ {
				scenicScore++
				if trees[i][rights] >= height {
					break
				}
			}
			scores = append(scores, scenicScore)

			// add all scores to the locations
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
