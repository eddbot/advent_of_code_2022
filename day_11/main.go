package main

import (
	"advent_of_code/parser"
	"advent_of_code/utils"
	"fmt"
	"sort"
	"strings"
)

type Monkey struct {
	items     []int64
	operation func(int64) int64
	test      func(int64) bool
	true      int
	false     int
	inspected int
}

func main() {
	input := parser.ReadInputFile(11)
	s := solution(input)
	fmt.Println(s)
}

func solution(input string) int {

	monkeySlice := []Monkey{}
	monkeys := strings.Split(input, "\n\n")
	var divisors = []int64{}
	for _, m := range monkeys {

		monkey := Monkey{items: []int64{}, inspected: 0}
		stats := strings.Split(m, "\n")

		// Monkey items
		ii := strings.Replace(stats[1], ",", " ", -1)
		items := strings.Split(ii, " ")[4:]
		for _, item := range items {
			if item == "" {
				continue
			}
			monkey.items = append(monkey.items, int64(utils.Conv(item)))
		}
		// operations
		oo := strings.Split(stats[2], "old ")
		o := strings.Split(oo[1], " ")
		monkey.operation = operationFactory(o[0], o[1])

		// tests
		tt := strings.Split(stats[3], " ")
		divisor := int64(utils.Conv(tt[len(tt)-1]))

		divisors = append(divisors, divisor)
		monkey.test = func(worry int64) bool {
			return worry%divisor == 0
		}

		// true
		tx := strings.Split(stats[4], " ")
		monkey.true = utils.Conv(tx[len(tx)-1])

		// false
		fx := strings.Split(stats[5], " ")
		monkey.false = utils.Conv(fx[len(fx)-1])

		monkeySlice = append(monkeySlice, monkey)

	}

	mod := divisors[0]

	for _, d := range divisors[1:] {
		mod *= d
	}

	// part 1 = 20
	// part 2 = 100000
	for i := 0; i < 20; i++ {
		for m := 0; m < len(monkeySlice); m++ {
			for _, item := range monkeySlice[m].items {
				// inspect
				monkeySlice[m].inspected++

				// manage worry levels

				// part 1
				x := monkeySlice[m].operation(item) / 3

				// part 2
				//x := monkeySlice[m].operation(item) % mod

				// test
				t := monkeySlice[m].test(x)
				// throw
				if t {
					monkeySlice[monkeySlice[m].true].items = append(monkeySlice[monkeySlice[m].true].items, x)
				} else {
					monkeySlice[monkeySlice[m].false].items = append(monkeySlice[monkeySlice[m].false].items, x)
				}

			}
			monkeySlice[m].items = []int64{}
			// we should have thrown all the items now
		}
	}

	m1, m2 := mostActiveMonkeys(monkeySlice)
	return m1 * m2
}

func mostActiveMonkeys(monkeys []Monkey) (int, int) {
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})

	return monkeys[0].inspected, monkeys[1].inspected
}

func operationFactory(operation, amount string) func(int64) int64 {

	var a int64
	if amount == "old" {
		operation = "square me ;)"
	} else {
		a = int64(utils.Conv(amount))
	}

	switch operation {
	case "*":
		return func(old int64) int64 {
			return old * a
		}
	case "/":
		return func(old int64) int64 {
			return old / a
		}
	case "-":
		return func(old int64) int64 {
			return old - a
		}
	case "+":
		return func(old int64) int64 {
			return old + a
		}
	default:
		return func(old int64) int64 {
			return old * old
		}
	}
}
