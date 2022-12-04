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
Benchmark/Day_01/Part_1-10  	   27750	     43257 ns/op	    4352 B/op	       7 allocs/op
Benchmark/Day_01/Part_2-10  	   27866	     42957 ns/op	    4352 B/op	       7 allocs/op
Benchmark/Day_02/Part_1-10  	   26811	     43803 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_02/Part_2-10  	   30726	     39068 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_03/Part_1-10  	   54710	     18781 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_03/Part_2-10  	   49051	     22843 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_04/Part_1-10  	   19765	     60946 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_04/Part_2-10  	   22245	     53727 ns/op	    4104 B/op	       2 allocs/op
PASS
ok  	github.com/armsnyder/aoc2022	12.803s
```
<!-- END BENCHMARKS -->
