package main

import (
	"bytes"
	"io"

	"github.com/armsnyder/aoc2022/aocutil"
)

var _ = declareDay(4, func(part2 bool, inputReader io.Reader) any {
	if part2 {
		return day04Part2(inputReader)
	}
	return day04Part1(inputReader)
})

func day04Part1(inputReader io.Reader) any {
	return day04WithCompareFunc(inputReader, func(a, b day04Range) bool {
		return a.contains(b) || b.contains(a)
	})
}

func day04Part2(inputReader io.Reader) any {
	return day04WithCompareFunc(inputReader, day04Range.overlaps)
}

func day04WithCompareFunc(inputReader io.Reader, compareFunc func(a, b day04Range) bool) int {
	total := 0

	aocutil.VisitStrings(inputReader, func(v []byte) {
		comma := bytes.IndexByte(v, ',')
		elf1 := day04ParseRange(v[:comma])
		elf2 := day04ParseRange(v[comma+1:])
		if compareFunc(elf1, elf2) {
			total++
		}
	})

	return total
}

type day04Range [2]int

func (a day04Range) contains(b day04Range) bool {
	return a[0] <= b[0] && a[1] >= b[1]
}

func (a day04Range) overlaps(b day04Range) bool {
	return (a[0] >= b[0] || a[1] >= b[0]) && (a[0] <= b[1] || a[1] <= b[1])
}

func day04ParseRange(v []byte) day04Range {
	hyphen := bytes.IndexByte(v, '-')
	min, _ := aocutil.AtoiBytes(v[:hyphen])
	max, _ := aocutil.AtoiBytes(v[hyphen+1:])
	return day04Range{min, max}
}
