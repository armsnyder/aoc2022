package main

import (
	"io"

	"github.com/armsnyder/aoc2022/aocutil"
)

var _ = declareDay(2, func(part2 bool, inputReader io.Reader) any {
	if part2 {
		return day02Part2(inputReader)
	}
	return day02Part1(inputReader)
})

func day02Part1(inputReader io.Reader) any {
	total := 0

	aocutil.VisitStrings(inputReader, func(v []byte) {
		p1 := int(v[0]) - 'A'
		p2 := int(v[2]) - 'X'
		switch (p2 - p1 + 3) % 3 {
		case 0:
			total += p2 + 4
		case 1:
			total += p2 + 7
		case 2:
			total += p2 + 1
		default:
			panic("can't get here")
		}
	})

	return total
}

func day02Part2(inputReader io.Reader) any {
	total := 0

	aocutil.VisitStrings(inputReader, func(v []byte) {
		p1 := int(v[0]) - 'A'
		switch v[2] {
		case 'X':
			total += (p1+2)%3 + 1
		case 'Y':
			total += p1 + 4
		case 'Z':
			total += (p1+1)%3 + 7
		default:
			panic(v[2])
		}
	})

	return total
}
