package main

import (
	"bytes"
	"io"
	"sort"

	"github.com/armsnyder/aoc2022/aocutil"
)

var _ = declareDay(13, func(part2 bool, inputReader io.Reader) any {
	if part2 {
		return day13Part2(inputReader)
	}
	return day13Part1(inputReader)
})

func day13Part1(inputReader io.Reader) any {
	result := 0
	pairIndex := 1

	aocutil.VisitStringGroupBytes(inputReader, func(v []byte) {
		newline := bytes.IndexByte(v, '\n')

		if day13Compare(v[1:newline-1], v[newline+2:len(v)-1]) == -1 {
			result += pairIndex
		}

		pairIndex++
	})

	return result
}

func day13Part2(inputReader io.Reader) any {
	sortedPackets := make([][]byte, 0, 300+2)
	sortedPackets = append(sortedPackets, []byte("[[2]]"))
	sortedPackets = append(sortedPackets, []byte("[[6]]"))
	marker2Index := 0
	marker6Index := 1

	aocutil.VisitStrings(inputReader, func(packet []byte) {
		insertIndex := sort.Search(len(sortedPackets), func(i int) bool {
			return day13Compare(packet, sortedPackets[i]) < 0
		})

		if insertIndex <= marker2Index {
			marker2Index++
		}

		if insertIndex <= marker6Index {
			marker6Index++
		}

		sortedPackets = append(sortedPackets, nil)
		copy(sortedPackets[insertIndex+1:], sortedPackets[insertIndex:])

		sortedPackets[insertIndex] = make([]byte, len(packet))
		copy(sortedPackets[insertIndex], packet)
	})

	return (marker2Index + 1) * (marker6Index + 1)
}

func day13Compare(left, right []byte) int {
	if len(left) == 0 {
		if len(right) == 0 {
			return 0
		}
		return -1
	}

	if len(right) == 0 {
		return 1
	}

	for {
		leftValueEnd := bytes.IndexByte(left, ',')
		if leftValueEnd == -1 {
			leftValueEnd = len(left)
		}

		rightValueEnd := bytes.IndexByte(right, ',')
		if rightValueEnd == -1 {
			rightValueEnd = len(right)
		}

		isLeftInt := left[0] != '['
		isRightInt := right[0] != '['

		switch {
		case !isLeftInt && !isRightInt:
			leftValue := left[1:day13FindClosingBracket(left[:])]
			rightValue := right[1:day13FindClosingBracket(right[:])]
			if result := day13Compare(leftValue, rightValue); result != 0 {
				return result
			}

		case isLeftInt && !isRightInt:
			leftValue := left[:leftValueEnd]
			rightValue := right[1:day13FindClosingBracket(right[:])]
			if result := day13Compare(leftValue, rightValue); result != 0 {
				return result
			}

		case !isLeftInt && isRightInt:
			leftValue := left[1:day13FindClosingBracket(left[:])]
			rightValue := right[:rightValueEnd]
			if result := day13Compare(leftValue, rightValue); result != 0 {
				return result
			}

		default:
			leftValue, _ := aocutil.AtoiBytes(left[:leftValueEnd])
			rightValue, _ := aocutil.AtoiBytes(right[:rightValueEnd])
			if leftValue < rightValue {
				return -1
			}
			if leftValue > rightValue {
				return 1
			}
		}

		if leftValueEnd == len(left) {
			if rightValueEnd == len(right) {
				return 0
			}
			return -1
		}

		if rightValueEnd == len(right) {
			return 1
		}

		left = left[leftValueEnd+1:]
		right = right[rightValueEnd+1:]
	}
}

func day13FindClosingBracket(s []byte) int {
	depth := 0
	for i, ch := range s {
		switch ch {
		case '[':
			depth++
		case ']':
			depth--
			if depth == 0 {
				return i
			}
		}
	}
	panic("no closing bracket")
}
