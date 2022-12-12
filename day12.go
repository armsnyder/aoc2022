package main

import (
	"io"
	"math"

	"github.com/armsnyder/aoc2022/aocutil"
)

var _ = declareDay(12, func(part2 bool, inputReader io.Reader) any {
	if part2 {
		return day12Part2(inputReader)
	}
	return day12Part1(inputReader)
})

func day12Part1(inputReader io.Reader) any {
	// This should have been a fun A* problem, but I decided to use a modified Dijkstra, since it is
	// orders of magnitude more efficient at this problem's scale.

	grid, width, start, end := day12Parse(inputReader)
	distances := make([]int, len(grid))
	for i := range distances {
		distances[i] = math.MaxInt
	}
	distances[start] = 0
	queue := make([]int, 0, len(grid))
	queue = append(queue, start)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, dir := range []int{1, -1, width, -width} {
			neighbor := cur + dir
			if neighbor < 0 || neighbor >= len(grid) {
				continue
			}
			if grid[cur]+1 < grid[neighbor] {
				continue
			}
			newDist := distances[cur] + 1
			oldDist := distances[neighbor]
			if newDist >= oldDist {
				continue
			}
			distances[neighbor] = newDist
			queue = append(queue, neighbor)
		}
	}

	return distances[end]
}

func day12Part2(inputReader io.Reader) any {
	grid, width, _, end := day12Parse(inputReader)
	distances := make([]int, len(grid))
	for i := range distances {
		distances[i] = math.MaxInt
	}
	distances[end] = 0
	queue := make([]int, 0, len(grid))
	queue = append(queue, end)
	minDistanceToA := math.MaxInt

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, dir := range []int{1, -1, width, -width} {
			neighbor := cur + dir
			if neighbor < 0 || neighbor >= len(grid) {
				continue
			}
			if grid[cur] > grid[neighbor]+1 {
				continue
			}
			newDist := distances[cur] + 1
			if newDist >= minDistanceToA {
				continue
			}
			oldDist := distances[neighbor]
			if newDist >= oldDist {
				continue
			}
			distances[neighbor] = newDist
			queue = append(queue, neighbor)
			if grid[neighbor] == 'a' {
				minDistanceToA = newDist
			}
		}
	}

	return minDistanceToA
}

func day12Parse(inputReader io.Reader) (grid []byte, width, start, end int) {
	var height int
	grid = make([]byte, 0, 154*41)
	aocutil.VisitStrings(inputReader, func(v []byte) {
		width = len(v)
		for i, ch := range v {
			switch ch {
			case 'S':
				start = width*height + i
				v[i] = 'a'
			case 'E':
				end = width*height + i
				v[i] = 'z'
			}
		}
		grid = append(grid, v...)
		height++
	})

	return grid, width, start, end
}
