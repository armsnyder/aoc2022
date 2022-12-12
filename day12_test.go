package main

import (
	"testing"
)

func TestDay12Part1(t *testing.T) {
	runDayTests(t, 12, []dayTest{
		{
			input: `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`,
			want: 31,
		},
	})
}

func TestDay12Part2(t *testing.T) {
	runDayTests(t, 12, []dayTest{
		{
			part2: true,
			input: `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`,
			want: 29,
		},
	})
}
