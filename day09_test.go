package main

import (
	"testing"
)

func TestDay09Part1(t *testing.T) {
	runDayTests(t, 9, []dayTest{
		{
			input: `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`,
			want: 13,
		},
	})
}

func TestDay09Part2(t *testing.T) {
	runDayTests(t, 9, []dayTest{
		{
			part2: true,
			input: `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`,
			want: 1,
		},
		{
			part2: true,
			input: `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`,
			want: 36,
		},
	})
}
