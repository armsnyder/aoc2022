package main

import (
	"testing"
)

func TestDay02Part1(t *testing.T) {
	runDayTests(t, 2, []dayTest{
		{
			input: `A Y
B X
C Z`,
			want: 15,
		},
	})
}

func TestDay02Part2(t *testing.T) {
	runDayTests(t, 2, []dayTest{
		{
			part2: true,
			input: `A Y
B X
C Z`,
			want: 12,
		},
	})
}
