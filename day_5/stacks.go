package main

import "advent_of_code/utils"

type stacks map[int][]string

func (s stacks) dup() stacks {
	newStack := stacks{}
	for k, v := range s {
		newStack[k] = v
	}
	return newStack
}

func (s stacks) remain(from, move int) []string {
	return s[from][move:]
}

func (s stacks) insert(to, from, index int) []string {
	return utils.Insert(s[to], s[from][index])
}
func (s stacks) insertMultiple(to, from, move int) []string {
	return utils.InsertMulti(s[to], s[from][0:move])
}
