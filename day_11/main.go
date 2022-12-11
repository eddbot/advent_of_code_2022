package main

import (
	"advent_of_code/parser"
	"advent_of_code/utils"
	"fmt"
	"sort"
	"strings"
)

type Monkey struct {
	items     []int
	operation func(int) int
	test      func(int) bool
	true      int
	false     int
	inspected int
}

func main() {
	//input := testInput()
	input := parser.ReadInputFile(11)
	p1 := part1(input)
	fmt.Println(p1)
}

func part1(input string) int {

	monkeySlice := []Monkey{}
	monkeys := strings.Split(input, "\n\n")

	for _, m := range monkeys {

		monkey := Monkey{items: []int{}, inspected: 0}
		stats := strings.Split(m, "\n")

		// Monkey items
		ii := strings.Replace(stats[1], ",", " ", -1)
		items := strings.Split(ii, " ")[4:]
		for _, item := range items {
			if item == "" {
				continue
			}
			monkey.items = append(monkey.items, utils.Conv(item))
		}
		// operations
		oo := strings.Split(stats[2], "old ")
		o := strings.Split(oo[1], " ")
		monkey.operation = operationFactory(o[0], o[1])

		// tests
		tt := strings.Split(stats[3], " ")
		divisor := utils.Conv(tt[len(tt)-1])
		monkey.test = func(worry int) bool {
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

	for i := 0; i < 10000; i++ {
		for m := 0; m < len(monkeySlice); m++ {
			for _, item := range monkeySlice[m].items {
				// inspect
				monkeySlice[m].inspected++
				x := monkeySlice[m].operation(item)
				// get bored
				x /= 3
				// test
				t := monkeySlice[m].test(x)
				// throw
				if t {
					monkeySlice[monkeySlice[m].true].items = append(monkeySlice[monkeySlice[m].true].items, x)
				} else {
					monkeySlice[monkeySlice[m].false].items = append(monkeySlice[monkeySlice[m].false].items, x)
				}

			}
			monkeySlice[m].items = []int{}
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

func operationFactory(operation, amount string) func(int) int {

	var a int
	if amount == "old" {
		operation = "square me bebbe"
	} else {
		a = utils.Conv(amount)
	}

	switch operation {
	case "*":
		return func(old int) int {
			return old * a
		}
	case "/":
		return func(old int) int {
			return old / a
		}
	case "-":
		return func(old int) int {
			return old - a
		}
	case "+":
		return func(old int) int {
			return old + a
		}
	default:
		return func(old int) int {
			return old * old
		}
	}
}

func testInput() string {
	return `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`
}
