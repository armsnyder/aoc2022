package main

import (
	"io"

	"github.com/armsnyder/aoc2022/aocutil"
)

var _ = declareDay(9, func(part2 bool, inputReader io.Reader) any {
	if part2 {
		return day09Knots(inputReader, 10)
	}
	return day09Knots(inputReader, 2)
})

func day09Knots(inputReader io.Reader, knots int) int {
	rope := make([]day09Coord, knots)
	visited := make(map[day09Coord]struct{})

	aocutil.VisitStrings(inputReader, func(v []byte) {
		dir := v[0]
		delta := getDelta(dir)
		steps, err := aocutil.AtoiBytes(v[2:])
		if err != nil {
			panic(err)
		}

		for i := 0; i < steps; i++ {
			rope[0] = rope[0].add(delta)

			for j := 1; j < knots; j++ {
				coord, moved := rope[j].follow(rope[j-1])
				if !moved {
					break
				}
				rope[j] = coord
			}

			visited[rope[knots-1]] = struct{}{}
		}
	})

	return len(visited)
}

type day09Coord [2]int

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func (c day09Coord) follow(c2 day09Coord) (day09Coord, bool) {
	dx := c2[0] - c[0]
	dy := c2[1] - c[1]

	if abs(dx) <= 1 && abs(dy) <= 1 {
		return c, false
	}

	if dx > 1 {
		dx = 1
	} else if dx < -1 {
		dx = -1
	}

	if dy > 1 {
		dy = 1
	} else if dy < -1 {
		dy = -1
	}

	c[0] += dx
	c[1] += dy

	return c, true
}

func (c day09Coord) add(c2 day09Coord) day09Coord {
	return day09Coord{c[0] + c2[0], c[1] + c2[1]}
}

func getDelta(dir byte) day09Coord {
	switch dir {
	case 'U':
		return day09Coord{0, 1}
	case 'D':
		return day09Coord{0, -1}
	case 'L':
		return day09Coord{-1, 0}
	case 'R':
		return day09Coord{1, 0}
	default:
		panic(dir)
	}
}
