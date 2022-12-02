# Advent of Code 2022 Solutions

My [Advent of Code 2022](https://adventofcode.com/2022) solutions.

Usage:

```
$ go run . [-d day] [-2]
```

Requires a `session.txt` file containing a session token, for pulling puzzle input and submitting
answers. (Inputs and answers are cached.)

## Benchmarks

Benchmarks use the real puzzle input, which is preloaded in memory.

<!-- BEGIN BENCHMARKS -->
```
goos: darwin
goarch: arm64
pkg: github.com/armsnyder/aoc2022
Benchmark/Day_01/Part_1-10  	   27234	     43117 ns/op	    4352 B/op	       7 allocs/op
Benchmark/Day_01/Part_2-10  	   27897	     42982 ns/op	    4352 B/op	       7 allocs/op
Benchmark/Day_02/Part_1-10  	   27349	     43862 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_02/Part_2-10  	   30909	     38919 ns/op	    4104 B/op	       2 allocs/op
PASS
ok  	github.com/armsnyder/aoc2022	6.630s
```
<!-- END BENCHMARKS -->
