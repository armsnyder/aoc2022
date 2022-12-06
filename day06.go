package main

import (
	"io"

	"github.com/armsnyder/aoc2022/aocutil"
)

var _ = declareDay(6, func(part2 bool, inputReader io.Reader) any {
	if part2 {
		return day06(inputReader, 14)
	}
	return day06(inputReader, 4)
})

func day06(inputReader io.Reader, packetMarkerSize int) any {
	var (
		streamIndex        int
		packetMarkerBuffer = make([]byte, packetMarkerSize)
	)

	aocutil.VisitCharactersWhile(inputReader, func(buf []byte) bool {
		for i := range buf {
			packetMarkerBuffer[streamIndex%packetMarkerSize] = buf[i] - 'a'

			if streamIndex >= packetMarkerSize && day06IsPacketMarker(packetMarkerBuffer) {
				return false
			}

			streamIndex++
		}

		return true
	})

	return streamIndex + 1
}

func day06IsPacketMarker(packet []byte) bool {
	var hits [26]bool

	for i := 0; i < len(packet); i++ {
		if hits[packet[i]] {
			return false
		}

		hits[packet[i]] = true
	}

	return true
}
