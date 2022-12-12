package main

import (
	"bytes"
	"io"

	"github.com/armsnyder/aoc2022/aocutil"
)

var _ = declareDay(10, func(part2 bool, inputReader io.Reader) any {
	if part2 {
		return day10Part2(inputReader)
	}
	return day10Part1(inputReader)
})

func day10Part1(inputReader io.Reader) any {
	sum := 0

	day10RunProgram(inputReader, func(cycle, x int) {
		if cycle == 19 || ((cycle-19)%40) == 0 {
			sum += ((cycle + 1) * x)
		}
	})

	return sum
}

func day10Part2(inputReader io.Reader) any {
	buf := bytes.NewBuffer(make([]byte, 0, 6*41+1))

	buf.WriteByte('\n')

	day10RunProgram(inputReader, func(cycle, x int) {
		head := cycle % 40
		spriteOffset := x - head

		if spriteOffset >= -1 && spriteOffset <= 1 {
			buf.WriteByte('#')
		} else {
			buf.WriteByte('.')
		}

		if head == 39 {
			buf.WriteByte('\n')
		}
	})

	return buf.String()
}

func day10RunProgram(inputReader io.Reader, tick func(cycle, x int)) {
	cycle := 0
	x := 1

	aocutil.VisitStrings(inputReader, func(v []byte) {
		tick(cycle, x)
		cycle++

		if string(v) == "noop" {
			return
		}

		tick(cycle, x)
		cycle++

		arg, err := aocutil.AtoiBytes(v[5:])
		if err != nil {
			panic(err)
		}

		x += arg
	})
}
