package main

import (
	"testing"
)

func TestDay01Part1(t *testing.T) {
	runDayTests(t, 1, []dayTest{
		{
			input: `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`,
			want: 24000,
		},
	})
}

func TestDay01Part2(t *testing.T) {
	runDayTests(t, 1, []dayTest{
		{
			part2: true,
			input: `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`,
			want: 45000,
		},
	})
}
