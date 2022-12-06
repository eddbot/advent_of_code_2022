package utils

import "strconv"

func Insert(in []string, val string) []string {
	temp := []string{val}
	temp = append(temp, in...)
	return temp
}

func InsertMulti(in []string, mul []string) []string {
	temp := []string{}
	temp = append(temp, mul...)
	temp = append(temp, in...)
	return temp
}

func Conv(in string) int {
	c, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return c
}
