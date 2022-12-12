package main

import (
	"io"
	"strconv"
	"strings"

	"github.com/armsnyder/aoc2022/aocutil"
	"github.com/samber/lo"
)

var _ = declareDay(11, func(part2 bool, inputReader io.Reader) any {
	if part2 {
		return day11Part2(inputReader)
	}
	return day11Part1(inputReader)
})

func day11Part1(inputReader io.Reader) any {
	return day11(inputReader, 20, func(_ []*day11Monkey) func(int) int {
		return func(worry int) int {
			return worry / 3
		}
	})
}

func day11Part2(inputReader io.Reader) any {
	return day11(inputReader, 10000, func(monkeys []*day11Monkey) func(int) int {
		modulus := lo.Reduce(monkeys, func(agg int, m *day11Monkey, _ int) int {
			return agg * m.divisor
		}, 1)

		return func(worry int) int {
			return worry % modulus
		}
	})
}

func day11(inputReader io.Reader, rounds int, makeWorryFunc func(monkeys []*day11Monkey) func(int) int) int {
	monkeys := day11Parse(inputReader)
	manageWorry := makeWorryFunc(monkeys)
	inspections := day11CountInspections(monkeys, rounds, manageWorry)
	monkeyBusiness := day11CalculateMonkeyBusiness(inspections)

	return monkeyBusiness
}

func day11CountInspections(monkeys []*day11Monkey, rounds int, manageWorry func(int) int) []int {
	inspections := make([]int, len(monkeys))

	for round := 0; round < rounds; round++ {
		for i, monkey := range monkeys {
			inspections[i] += len(monkey.items)

			for _, item := range monkey.items {
				item = monkey.operation(item)
				item = manageWorry(item)

				if item%monkey.divisor == 0 {
					monkeys[monkey.ifTrue].items = append(monkeys[monkey.ifTrue].items, item)
				} else {
					monkeys[monkey.ifFalse].items = append(monkeys[monkey.ifFalse].items, item)
				}
			}

			monkey.items = monkey.items[:0]
		}
	}

	return inspections
}

func day11CalculateMonkeyBusiness(inspections []int) int {
	var top2 [2]int

	for _, v := range inspections {
		switch {
		case v > top2[1]:
			top2[0], top2[1] = top2[1], v
		case v > top2[0]:
			top2[0] = v
		}
	}

	return top2[0] * top2[1]
}

type day11Monkey struct {
	items     []int
	operation func(int) int
	divisor   int
	ifTrue    int
	ifFalse   int
}

func day11Parse(inputReader io.Reader) []*day11Monkey {
	var monkeys []*day11Monkey

	aocutil.VisitStringGroups(inputReader, func(v []string) {
		var monkey day11Monkey

		monkey.items = lo.Map(strings.Split(v[1][strings.LastIndexByte(v[1], ':')+2:], ", "), func(s string, _ int) int {
			item, _ := strconv.Atoi(s)
			return item
		})

		lastSpace := strings.LastIndexByte(v[2], ' ')
		op := v[2][lastSpace-1]
		operand := v[2][lastSpace+1:]

		var opFn func(a, b int) int

		switch op {
		case '+':
			opFn = func(a, b int) int {
				return a + b
			}
		case '*':
			opFn = func(a, b int) int {
				return a * b
			}
		}

		switch operand {
		case "old":
			monkey.operation = func(i int) int {
				return opFn(i, i)
			}
		default:
			operandInt, _ := strconv.Atoi(operand)
			monkey.operation = func(i int) int {
				return opFn(i, operandInt)
			}
		}

		monkey.divisor, _ = strconv.Atoi(v[3][strings.LastIndexByte(v[3], ' ')+1:])
		monkey.ifTrue, _ = strconv.Atoi(v[4][strings.LastIndexByte(v[4], ' ')+1:])
		monkey.ifFalse, _ = strconv.Atoi(v[5][strings.LastIndexByte(v[5], ' ')+1:])

		monkeys = append(monkeys, &monkey)
	})

	return monkeys
}
