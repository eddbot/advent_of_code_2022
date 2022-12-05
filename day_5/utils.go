package main

import "strconv"

func insert(in []string, val string) []string {
	temp := []string{val}
	temp = append(temp, in...)
	return temp
}

func insertMulti(in []string, mul []string) []string {
	temp := []string{}
	temp = append(temp, mul...)
	temp = append(temp, in...)
	return temp
}

func conv(in string) int {
	c, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return c
}

func answerParser(stacks map[int][]string) string {
	answer := ""
	for i := 0; i < len(stacks); i++ {
		answer += stacks[i+1][0]
	}
	return answer
}
