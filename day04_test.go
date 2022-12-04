package main

import (
	"testing"
)

func TestDay04Part1(t *testing.T) {
	runDayTests(t, 4, []dayTest{
		{
			input: `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`,
			want: 2,
		},
	})
}

func TestDay04Part2(t *testing.T) {
	runDayTests(t, 4, []dayTest{
		{
			part2: true,
			input: `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`,
			want: 4,
		},
	})
}
