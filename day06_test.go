package main

import (
	"testing"
)

func TestDay06Part1(t *testing.T) {
	runDayTests(t, 6, []dayTest{
		{
			input: `mjqjpqmgbljsphdztnvjfqwrcgsmlb`,
			want:  7,
		},
		{
			input: `bvwbjplbgvbhsrlpgdmjqwftvncz`,
			want:  5,
		},
		{
			input: `nppdvjthqldpwncqszvftbrmjlhg`,
			want:  6,
		},
		{
			input: `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`,
			want:  10,
		},
		{
			input: `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`,
			want:  11,
		},
	})
}

func TestDay06Part2(t *testing.T) {
	runDayTests(t, 6, []dayTest{
		{
			part2: true,
			input: `mjqjpqmgbljsphdztnvjfqwrcgsmlb`,
			want:  19,
		},
		{
			part2: true,
			input: `bvwbjplbgvbhsrlpgdmjqwftvncz`,
			want:  23,
		},
		{
			part2: true,
			input: `nppdvjthqldpwncqszvftbrmjlhg`,
			want:  23,
		},
		{
			part2: true,
			input: `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`,
			want:  29,
		},
		{
			part2: true,
			input: `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`,
			want:  26,
		},
	})
}
