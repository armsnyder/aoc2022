package main

import (
	"io"
	"sort"

	"github.com/armsnyder/aoc2022/aocutil"
	"github.com/samber/lo"
)

var _ = declareDay(1, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day01Part2(inputReader)
	}
	return day01Part1(inputReader)
})

func day01Part1(inputReader io.Reader) interface{} {
	max := 0

	aocutil.VisitIntGroups(inputReader, func(v []int) {
		max = lo.Max([]int{max, lo.Sum(v)})
	})

	return max
}

func day01Part2(inputReader io.Reader) interface{} {
	var top3 [3]int

	aocutil.VisitIntGroups(inputReader, func(v []int) {
		sum := lo.Sum(v)
		if sum <= top3[0] {
			return
		}
		i := sort.SearchInts(top3[:], sum)
		copy(top3[:], top3[1:i])
		top3[i-1] = sum
	})

	return lo.Sum(top3[:])
}
