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

		left, _ := day13ParsePacket(v[:newline])
		right, _ := day13ParsePacket(v[newline+1:])

		if day13ComparePackets(left, right) == -1 {
			result += pairIndex
		}

		pairIndex++
	})

	return result
}

func day13Part2(inputReader io.Reader) any {
	sortedPackets := make([]day13Packet, 0, 300+2)
	marker2, _ := day13ParsePacket([]byte("[[2]]"))
	marker6, _ := day13ParsePacket([]byte("[[6]]"))
	sortedPackets = append(sortedPackets, marker2, marker6)
	marker2Index, marker6Index := 0, 1

	aocutil.VisitStrings(inputReader, func(b []byte) {
		// Parse the packet.
		packet, _ := day13ParsePacket(b)

		// Find the index where we should insert the packet to preserve sort order.
		insertIndex := sort.Search(len(sortedPackets), func(i int) bool {
			return day13ComparePackets(packet, sortedPackets[i]) < 0
		})

		// Track the marker indices if they are about to be shifted by the insert.
		if insertIndex <= marker2Index {
			marker2Index++
		}
		if insertIndex <= marker6Index {
			marker6Index++
		}

		// Insert the packet at the correct index to preserve sort order.
		sortedPackets = append(sortedPackets, day13Packet{})
		copy(sortedPackets[insertIndex+1:], sortedPackets[insertIndex:])
		sortedPackets[insertIndex] = packet
	})

	return (marker2Index + 1) * (marker6Index + 1)
}

// day13Packet can be either an integer or list. If it is a list, then integer is set to -1.
type day13Packet struct {
	integer int
	list    []day13Packet
}

// day13ParsePacket parses the packet and returns the number of bytes read.
func day13ParsePacket(b []byte) (day13Packet, int) {
	switch b[0] {
	case ']':
		// Case: Reached end of packet.
		return day13Packet{}, 0

	case '[':
		// Case: Packet is a list. Iterate over packet values.
		packet := day13Packet{integer: -1}
		cursor := 1
		for b[cursor-1] != ']' {
			// Parse each item as a packet, recusively.
			child, bytesRead := day13ParsePacket(b[cursor:])
			if bytesRead != 0 {
				packet.list = append(packet.list, child)
			}
			cursor += bytesRead + 1
		}
		return packet, cursor

	default:
		// Case: Packet is an integer.
		valueEnd := bytes.IndexAny(b, ",]")
		value, _ := aocutil.AtoiBytes(b[:valueEnd])
		return day13Packet{integer: value}, valueEnd
	}
}

// day13ComparePackets returns -1 if left < right, 1 if left > right, or 0 if they are equal.
func day13ComparePackets(left, right day13Packet) int {
	switch {
	case left.integer >= 0 && right.integer >= 0:
		// Case: Both are integers.
		if left.integer > right.integer {
			return 1
		}
		if left.integer < right.integer {
			return -1
		}
		return 0

	case left.integer >= 0 && right.integer == -1:
		// Case: Left is an integer and right is a list. Recurse to compare both as lists.
		return day13ComparePackets(day13Packet{integer: -1, list: []day13Packet{left}}, right)

	case right.integer >= 0 && left.integer == -1:
		// Case: Right is an integer and left is a list. Recurse to compare both as lists.
		return day13ComparePackets(left, day13Packet{integer: -1, list: []day13Packet{right}})

	default:
		// Case: Both left and right are lists. Compare individual values.
		for i := 0; ; i++ {
			if i == len(left.list) {
				if i == len(right.list) {
					// Both left and right ran out of values.
					return 0
				}
				// Left ran out of values first.
				return -1
			}
			if i == len(right.list) {
				// Right ran out of values first.
				return 1
			}
			if result := day13ComparePackets(left.list[i], right.list[i]); result != 0 {
				// Comparing both values gave us a result. We're done.
				return result
			}
		}
	}
}
