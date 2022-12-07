package main

import (
	"bytes"
	"io"
	"sort"

	"github.com/armsnyder/aoc2022/aocutil"
	"github.com/samber/lo"
)

var _ = declareDay(7, func(part2 bool, inputReader io.Reader) any {
	if part2 {
		return day07Part2(inputReader)
	}
	return day07Part1(inputReader)
})

func day07Part1(inputReader io.Reader) any {
	root := day07BuildFileSystem(inputReader)
	total := 0

	var visit func(node *day07File) int
	visit = func(node *day07File) int {
		if len(node.children) == 0 {
			return node.size
		}

		size := 0

		for _, child := range node.children {
			size += visit(child)
		}

		if size <= 100000 {
			total += size
		}

		return size
	}

	visit(root)

	return total
}

func day07Part2(inputReader io.Reader) any {
	root := day07BuildFileSystem(inputReader)
	var sortedDirSizes []int

	var visit func(node *day07File) int
	visit = func(node *day07File) int {
		if len(node.children) == 0 {
			return node.size
		}

		size := 0

		for _, child := range node.children {
			size += visit(child)
		}

		insertIndex := sort.SearchInts(sortedDirSizes, size)
		sortedDirSizes = append(sortedDirSizes, 0)
		copy(sortedDirSizes[insertIndex+1:], sortedDirSizes[insertIndex:])
		sortedDirSizes[insertIndex] = size

		return size
	}

	const diskSize = 70000000
	const needUnusedSpace = 30000000
	fileSystemSize := visit(root)
	needDeleted := needUnusedSpace - (diskSize - fileSystemSize)
	minSizeIndex := sort.SearchInts(sortedDirSizes, needDeleted)

	return sortedDirSizes[minSizeIndex]
}

type day07File struct {
	size     int
	name     string
	parent   *day07File
	children []*day07File
}

func day07BuildFileSystem(inputReader io.Reader) *day07File {
	root := &day07File{name: "/"}
	var cur *day07File

	aocutil.VisitStrings(inputReader, func(v []byte) {
		if v[0] == '$' {
			if string(v[2:4]) != "cd" {
				return
			}

			cdDir := string(v[bytes.LastIndexByte(v, ' ')+1:])

			switch cdDir {
			case "/":
				cur = root
			case "..":
				cur = cur.parent
			default:
				cur, _ = lo.Find(cur.children, func(child *day07File) bool {
					return child.name == cdDir
				})
			}

			return
		}

		spaceIndex := bytes.IndexByte(v, ' ')
		name := string(v[spaceIndex+1:])

		if !lo.ContainsBy(cur.children, func(node *day07File) bool {
			return node.name == name
		}) {
			newNode := &day07File{
				name:   name,
				parent: cur,
			}
			if string(v[:spaceIndex]) != "dir" {
				newNode.size, _ = aocutil.AtoiBytes(v[:spaceIndex])
			}
			cur.children = append(cur.children, newNode)
		}
	})

	return root
}
