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
Benchmark/Day_01/Part_1-10  	   26278	     45972 ns/op	    4352 B/op	       7 allocs/op
Benchmark/Day_01/Part_2-10  	   26640	     44977 ns/op	    4352 B/op	       7 allocs/op
Benchmark/Day_02/Part_1-10  	   26178	     46034 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_02/Part_2-10  	   28489	     42024 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_03/Part_1-10  	   84398	     14252 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_03/Part_2-10  	   71409	     14897 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_04/Part_1-10  	   19555	     61567 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_04/Part_2-10  	   22539	     53138 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_05/Part_1-10  	   16692	     73375 ns/op	   77704 B/op	     538 allocs/op
Benchmark/Day_05/Part_2-10  	   17144	     69007 ns/op	   77728 B/op	     538 allocs/op
Benchmark/Day_06/Part_1-10  	  152394	      7527 ns/op	     528 B/op	       3 allocs/op
Benchmark/Day_06/Part_2-10  	  105369	     11281 ns/op	     536 B/op	       3 allocs/op
PASS
ok  	github.com/armsnyder/aoc2022	19.265s
```
<!-- END BENCHMARKS -->
