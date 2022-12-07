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
Benchmark/Day_01/Part_1-10  	   25965	     46407 ns/op	    4352 B/op	       7 allocs/op
Benchmark/Day_01/Part_2-10  	   26155	     45665 ns/op	    4352 B/op	       7 allocs/op
Benchmark/Day_02/Part_1-10  	   25891	     46279 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_02/Part_2-10  	   28380	     42458 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_03/Part_1-10  	   89030	     14454 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_03/Part_2-10  	   72481	     16248 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_04/Part_1-10  	   19263	     62516 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_04/Part_2-10  	   21831	     54195 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_05/Part_1-10  	   16362	     73172 ns/op	   77704 B/op	     538 allocs/op
Benchmark/Day_05/Part_2-10  	   16782	     71921 ns/op	   77728 B/op	     538 allocs/op
Benchmark/Day_06/Part_1-10  	  151689	      7595 ns/op	     528 B/op	       3 allocs/op
Benchmark/Day_06/Part_2-10  	  104497	     11431 ns/op	     536 B/op	       3 allocs/op
Benchmark/Day_07/Part_1-10  	   19945	     59614 ns/op	   44472 B/op	    1268 allocs/op
Benchmark/Day_07/Part_2-10  	   17802	     66256 ns/op	   48552 B/op	    1277 allocs/op
PASS
ok  	github.com/armsnyder/aoc2022	23.179s
```
<!-- END BENCHMARKS -->
