package main

import (
	"container/heap"
	"fmt"
	"io"
	"math"

	"github.com/armsnyder/aoc2022/aocutil"
)

var _ = declareDay(12, func(part2 bool, inputReader io.Reader) any {
	if part2 {
		return day12Part2(inputReader)
	}
	return day12Part12(inputReader)
})

func day12Part12(inputReader io.Reader) any {
	loops := 0

	grid := aocutil.Read2DByteArray(inputReader)

	var start day12Coord
	var end day12Coord

	// TODO optimize break
	for i := range grid {
		for j := range grid[i] {
			switch grid[i][j] {
			case 'S':
				grid[i][j] = 'a'
				start = day12Coord{i, j}
			case 'E':
				grid[i][j] = 'z'
				end = day12Coord{i, j}
			}
		}
	}

	// A* search algorithm: https://en.wikipedia.org/wiki/A*_search_algorithm

	h := func(coord day12Coord) int {
		di := end[0] - coord[0]
		if di < 0 {
			di *= -1
		}
		dj := end[1] - coord[1]
		if dj < 0 {
			dj *= -1
		}
		manhattanDist := di + dj
		heightDist := int('z' - grid[coord[0]][coord[1]])
		if heightDist > manhattanDist {
			return heightDist
		}
		return manhattanDist
	}

	openSet := &day12MinHeap{}
	heap.Push(openSet, &day12MinHeapItem{coord: start, fScore: h(start)})
	closedSet := make(map[day12Coord]int)
	gScore := make(map[day12Coord]int)
	gScore[start] = 0

	for openSet.Len() > 0 {
		current := heap.Pop(openSet).(*day12MinHeapItem)
		loops++
		if current.coord == end {
			fmt.Printf("loops: %d\n", loops)
			return current.fScore
		}
		for _, direction := range []day12Coord{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			neighbor := day12Coord{current.coord[0] + direction[0], current.coord[1] + direction[1]}

			if neighbor[0] < 0 || neighbor[0] >= len(grid) || neighbor[1] < 0 || neighbor[1] >= len(grid[0]) {
				continue
			}

			currentHeight := grid[current.coord[0]][current.coord[1]]
			neighborHeight := grid[neighbor[0]][neighbor[1]]
			if neighborHeight > currentHeight+1 {
				continue
			}

			currentGScore, exists := gScore[current.coord]
			if !exists {
				currentGScore = math.MaxInt
			}

			neighborGScore, exists := gScore[neighbor]
			if !exists {
				neighborGScore = math.MaxInt
			}

			tentativeGScore := currentGScore + 1

			if tentativeGScore < neighborGScore {
				gScore[neighbor] = tentativeGScore
				openIndex, openFScore := openSet.GetFScore(neighbor)
				fScore := tentativeGScore + h(neighbor)
				if openFScore < fScore {
					continue
				}
				closedFScore, exists := closedSet[neighbor]
				if exists && closedFScore < fScore {
					continue
				}
				if exists {
					fmt.Println("revisiting")
				}
				if openIndex > -1 {
					heap.Remove(openSet, openIndex)
				}
				heap.Push(openSet, &day12MinHeapItem{
					coord:  neighbor,
					fScore: fScore,
				})

				// if !lo.ContainsBy(openSet, func(item day12MinHeapItem) bool {
				// 	return item.coord == neighbor
				// }) {
				// 	heap.Push(&openSet, day12MinHeapItem{
				// 		coord:  neighbor,
				// 		fScore: tentativeGScore + h(neighbor),
				// 	})
				// }
			}
		}
		closedSet[current.coord] = current.fScore
	}

	panic("no path")
}

type day12MinHeapItem struct {
	coord  day12Coord
	fScore int
}

type day12MinHeap struct {
	values []*day12MinHeapItem
	index  map[day12Coord]int
}

func (h day12MinHeap) Len() int {
	return len(h.values)
}

func (h day12MinHeap) Less(i, j int) bool {
	return h.values[i].fScore < h.values[j].fScore
}

func (h day12MinHeap) Swap(i, j int) {
	iCoord := h.values[i].coord
	jCoord := h.values[j].coord
	h.index[iCoord], h.index[jCoord] = h.index[jCoord], h.index[iCoord]
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

func (h *day12MinHeap) Push(x any) {
	item := x.(*day12MinHeapItem)
	if i, exists := h.index[item.coord]; exists {
		panic(fmt.Errorf("%#v already exists as %#v", item, h.values[i]))
	}
	if h.index == nil {
		h.index = make(map[day12Coord]int)
	}
	h.index[item.coord] = len(h.values)
	h.values = append(h.values, item)
}

func (h *day12MinHeap) Pop() any {
	old := h.values
	n := len(old) - 1
	x := old[n]
	h.values = old[:n]
	delete(h.index, x.coord)
	return x
}

func (h day12MinHeap) GetFScore(coord day12Coord) (index, fScore int) {
	if i, ok := h.index[coord]; ok {
		return i, h.values[i].fScore
	}
	return -1, math.MaxInt
}

func day12Part1(inputReader io.Reader) any {
	loops := 0

	grid := aocutil.Read2DByteArray(inputReader)
	var start day12Coord
	var end day12Coord
	for i := range grid {
		for j := range grid[i] {
			switch grid[i][j] {
			case 'S':
				start = day12Coord{i, j}
			case 'E':
				end = day12Coord{i, j}
			}
		}
	}

	stack := []day12Item{{coord: start}}

	visited := make(map[day12Coord]int)
	visited[start] = 0

	visit := func(cur day12Coord, item day12Item) {
		if item.coord[0] < 0 || item.coord[0] >= len(grid) || item.coord[1] < 0 || item.coord[1] >= len(grid[0]) {
			return
		}
		curCh := grid[cur[0]][cur[1]]
		itemCh := grid[item.coord[0]][item.coord[1]]
		if itemCh == 'E' {
			itemCh = 'z'
		}
		if itemCh == 'S' {
			itemCh = 'a'
		}
		if curCh == 'E' {
			curCh = 'z'
		}
		if curCh == 'S' {
			curCh = 'a'
		}
		if curCh+1 < itemCh {
			return
		}
		oldCost, exists := visited[item.coord]
		if exists && item.cost >= oldCost {
			return
		}
		visited[item.coord] = item.cost
		stack = append(stack, item)
	}

	for len(stack) > 0 {
		loops++
		cur := stack[0]
		stack = stack[1:]

		visit(cur.coord, day12Item{cost: cur.cost + 1, coord: day12Coord{cur.coord[0] - 1, cur.coord[1]}})
		visit(cur.coord, day12Item{cost: cur.cost + 1, coord: day12Coord{cur.coord[0] + 1, cur.coord[1]}})
		visit(cur.coord, day12Item{cost: cur.cost + 1, coord: day12Coord{cur.coord[0], cur.coord[1] - 1}})
		visit(cur.coord, day12Item{cost: cur.cost + 1, coord: day12Coord{cur.coord[0], cur.coord[1] + 1}})
	}

	fmt.Printf("loops: %d\n", loops)
	return visited[end]
}

func day12Part2(inputReader io.Reader) any {
	grid := aocutil.Read2DByteArray(inputReader)
	var end day12Coord
	for i := range grid {
		for j := range grid[i] {
			switch grid[i][j] {
			case 'E':
				end = day12Coord{i, j}
			}
		}
	}

	stack := []day12Item{{coord: end}}

	visited := make(map[day12Coord]int)
	visited[end] = 0

	minA := math.MaxInt

	visit := func(cur day12Coord, item day12Item) {
		if item.coord[0] < 0 || item.coord[0] >= len(grid) || item.coord[1] < 0 || item.coord[1] >= len(grid[0]) {
			return
		}
		curCh := grid[cur[0]][cur[1]]
		itemCh := grid[item.coord[0]][item.coord[1]]
		if itemCh == 'E' {
			itemCh = 'z'
		}
		if itemCh == 'S' {
			itemCh = 'a'
		}
		if curCh == 'E' {
			curCh = 'z'
		}
		if curCh == 'S' {
			curCh = 'a'
		}
		if curCh > itemCh+1 {
			return
		}
		oldCost, exists := visited[item.coord]
		if exists && item.cost >= oldCost {
			return
		}
		visited[item.coord] = item.cost
		stack = append(stack, item)
		if itemCh == 'a' && item.cost < minA {
			minA = item.cost
		}
	}

	for len(stack) > 0 {
		cur := stack[0]
		stack = stack[1:]

		visit(cur.coord, day12Item{cost: cur.cost + 1, coord: day12Coord{cur.coord[0] - 1, cur.coord[1]}})
		visit(cur.coord, day12Item{cost: cur.cost + 1, coord: day12Coord{cur.coord[0] + 1, cur.coord[1]}})
		visit(cur.coord, day12Item{cost: cur.cost + 1, coord: day12Coord{cur.coord[0], cur.coord[1] - 1}})
		visit(cur.coord, day12Item{cost: cur.cost + 1, coord: day12Coord{cur.coord[0], cur.coord[1] + 1}})
	}

	return minA
}

type day12Coord [2]int

type day12Item struct {
	coord day12Coord
	cost  int
}
