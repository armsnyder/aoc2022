package main

import (
	"bytes"
	"io"

	"github.com/armsnyder/aoc2022/aocutil"
	"github.com/samber/lo"
)

var _ = declareDay(5, func(part2 bool, inputReader io.Reader) any {
	if part2 {
		return day05Part2(inputReader)
	}
	return day05Part1(inputReader)
})

func day05Part1(inputReader io.Reader) any {
	return day05UsingCrane(inputReader, func(stacks *day05Stacks, n, from, to int) {
		for i := 0; i < n; i++ {
			(*stacks)[to] = append((*stacks)[to], (*stacks)[from][len((*stacks)[from])-1])
			(*stacks)[from] = (*stacks)[from][:len((*stacks)[from])-1]
		}
	})
}

func day05Part2(inputReader io.Reader) any {
	return day05UsingCrane(inputReader, func(stacks *day05Stacks, n, from, to int) {
		(*stacks)[to] = append((*stacks)[to], (*stacks)[from][len((*stacks)[from])-n:]...)
		(*stacks)[from] = (*stacks)[from][:len((*stacks)[from])-n]
	})
}

type day05Stacks [9][]byte

func day05UsingCrane(inputReader io.Reader, crane func(stacks *day05Stacks, n, from, to int)) string {
	var stacks day05Stacks
	scanningCrates := true

	aocutil.VisitStrings(inputReader, func(v []byte) {
		if v[1] > '0' && v[1] <= '9' {
			scanningCrates = false
			for i := range stacks {
				lo.Reverse(stacks[i])
			}
			return
		}

		if scanningCrates {
			for crateIndex := 1; crateIndex < len(v); crateIndex += 4 {
				stackIndex := crateIndex / 4
				if v[crateIndex] != ' ' {
					stacks[stackIndex] = append(stacks[stackIndex], v[crateIndex])
				}
			}
			return
		}

		split := bytes.SplitN(v, []byte{' '}, 6)
		n, _ := aocutil.AtoiBytes(split[1])
		from, _ := aocutil.AtoiBytes(split[3])
		from -= 1
		to, _ := aocutil.AtoiBytes(split[5])
		to -= 1
		crane(&stacks, n, from, to)
	})

	result := make([]byte, 0, len(stacks))
	for _, stack := range stacks {
		if len(stack) == 0 {
			break
		}
		result = append(result, stack[len(stack)-1])
	}

	return string(result)
}
