package main

import (
	"io"

	"github.com/armsnyder/aoc2022/aocutil"
)

var _ = declareDay(3, func(part2 bool, inputReader io.Reader) any {
	if part2 {
		return day03Part2(inputReader)
	}
	return day03Part1(inputReader)
})

func day03Part1(inputReader io.Reader) any {
	var cache [26*2 + 1]bool
	total := 0

	aocutil.VisitStrings(inputReader, func(v []byte) {
		for _, b := range v[:len(v)/2] {
			p := day03Priority(b)
			cache[p] = true
		}

		for _, b := range v[len(v)/2:] {
			p := day03Priority(b)
			if cache[p] {
				total += p
				break
			}
		}

		for i := range cache {
			cache[i] = false
		}
	})

	return total
}

func day03Part2(inputReader io.Reader) any {
	var cache [26*2 + 1]int
	total := 0
	lineIndex := 0

	aocutil.VisitStrings(inputReader, func(v []byte) {
		defer func() {
			lineIndex++
		}()

		elfIndex := lineIndex % 3

		for _, b := range v {
			p := day03Priority(b)

			if cache[p] == elfIndex {
				if elfIndex == 2 {
					total += p

					for i := range cache {
						cache[i] = 0
					}

					return
				}

				cache[p] = elfIndex + 1
			}
		}
	})

	return total
}

func day03Priority(item byte) int {
	if item >= 'a' {
		return int(item) - int('a') + 1
	}
	return int(item) - int('A') + 26 + 1
}
