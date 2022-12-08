package main

import (
	"testing"
)

func TestDay08Part1(t *testing.T) {
	runDayTests(t, 8, []dayTest{
		{
			input: `30373
25512
65332
33549
35390`,
			want: 21,
		},
	})
}

func TestDay08Part2(t *testing.T) {
	runDayTests(t, 8, []dayTest{
		{
			part2: true,
			input: `30373
25512
65332
33549
35390`,
			want: 8,
		},
	})
}
