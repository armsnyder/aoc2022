package main

import (
	"bufio"
	"io"
)

var _ = declareDay(8, func(part2 bool, inputReader io.Reader) any {
	if part2 {
		return day08Part2(inputReader)
	}
	return day08Part1(inputReader)
})

func day08Part1(inputReader io.Reader) any {
	grid, size := day08ParseGrid(inputReader)
	visible := make([]bool, len(grid))
	total := 0

	addVisible := func(j int) {
		if !visible[j] {
			total++
			visible[j] = true
		}
	}

	for i := 0; i < size; i++ {
		// Left to right.
		var max byte
		for j := i * size; j < (i+1)*size; j++ {
			if grid[j] > max {
				max = grid[j]
				addVisible(j)
			}
		}

		// Right to left.
		max = 0
		for j := i*size + size - 1; j >= i*size; j-- {
			if grid[j] > max {
				max = grid[j]
				addVisible(j)
			}
		}

		// Top to bottom.
		max = 0
		for j := i; j < len(grid); j += size {
			if grid[j] > max {
				max = grid[j]
				addVisible(j)
			}
		}

		// Bottom to top.
		max = 0
		for j := size*(size-1) + i; j >= 0; j -= size {
			if grid[j] > max {
				max = grid[j]
				addVisible(j)
			}
		}
	}

	return total
}

func day08Part2(inputReader io.Reader) any {
	grid, size := day08ParseGrid(inputReader)
	highScore := 0
	for row := 1; row < size-1; row++ {
		for col := 1; col < size-1; col++ {
			if score := day08Score(grid, size, row, col); score > highScore {
				highScore = score
			}
		}
	}
	return highScore
}

func day08Score(grid []byte, size, row, col int) int {
	treeIndex := row*size + col
	treeHeight := grid[treeIndex]
	scenicScore := 1

	// Looking right.
	viewingDistance := 1
	for i := treeIndex + 1; i < (row+1)*size-1; i++ {
		if grid[i] >= treeHeight {
			break
		}
		viewingDistance++
	}
	scenicScore *= viewingDistance

	// Looking left
	viewingDistance = 1
	for i := treeIndex - 1; i > row*size; i-- {
		if grid[i] >= treeHeight {
			break
		}
		viewingDistance++
	}
	scenicScore *= viewingDistance

	// Looking down.
	viewingDistance = 1
	for i := treeIndex + size; i < len(grid)-size; i += size {
		if grid[i] >= treeHeight {
			break
		}
		viewingDistance++
	}
	scenicScore *= viewingDistance

	// Looking up.
	viewingDistance = 1
	for i := treeIndex - size; i >= size; i -= size {
		if grid[i] >= treeHeight {
			break
		}
		viewingDistance++
	}
	scenicScore *= viewingDistance

	return scenicScore
}

func day08ParseGrid(inputReader io.Reader) (grid []byte, size int) {
	scanner := bufio.NewScanner(inputReader)

	for i := 0; scanner.Scan(); i++ {
		if size == 0 {
			size = len(scanner.Bytes())
			grid = make([]byte, size*size)
		}
		copy(grid[i*size:], scanner.Bytes())
	}

	return grid, size
}
