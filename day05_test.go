package main

import (
	"testing"
)

func TestDay05Part1(t *testing.T) {
	runDayTests(t, 5, []dayTest{
		{
			input: `    [D]    
[N] [C]    
[Z] [M] [P]
	1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`,
			want: "CMZ",
		},
	})
}

func TestDay05Part2(t *testing.T) {
	runDayTests(t, 5, []dayTest{
		{
			part2: true,
			input: `    [D]    
[N] [C]    
[Z] [M] [P]
	1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`,
			want: "MCD",
		},
	})
}
