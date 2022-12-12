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
Benchmark/Day_01/Part_1-10  	   26448	     45013 ns/op	    4352 B/op	       7 allocs/op
Benchmark/Day_01/Part_2-10  	   26694	     44850 ns/op	    4352 B/op	       7 allocs/op
Benchmark/Day_02/Part_1-10  	   26480	     45248 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_02/Part_2-10  	   28912	     41424 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_03/Part_1-10  	   61839	     16699 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_03/Part_2-10  	   49460	     20532 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_04/Part_1-10  	   19546	     61466 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_04/Part_2-10  	   22556	     53127 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_05/Part_1-10  	   16776	     71551 ns/op	   77704 B/op	     538 allocs/op
Benchmark/Day_05/Part_2-10  	   17480	     68590 ns/op	   77728 B/op	     538 allocs/op
Benchmark/Day_06/Part_1-10  	  153640	      7526 ns/op	     528 B/op	       3 allocs/op
Benchmark/Day_06/Part_2-10  	  106040	     11126 ns/op	     536 B/op	       3 allocs/op
Benchmark/Day_07/Part_1-10  	   20505	     58892 ns/op	   44472 B/op	    1268 allocs/op
Benchmark/Day_07/Part_2-10  	   18646	     64747 ns/op	   48552 B/op	    1277 allocs/op
Benchmark/Day_08/Part_1-10  	   31138	     38112 ns/op	   24584 B/op	       4 allocs/op
Benchmark/Day_08/Part_2-10  	    5697	    192334 ns/op	   14344 B/op	       3 allocs/op
Benchmark/Day_09/Part_1-10  	    1828	    646449 ns/op	  357942 B/op	     191 allocs/op
Benchmark/Day_09/Part_2-10  	    1107	   1079871 ns/op	  177545 B/op	     101 allocs/op
Benchmark/Day_10/Part_1-10  	  374060	      3115 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_10/Part_2-10  	  359282	      3270 ns/op	    4368 B/op	       3 allocs/op
Benchmark/Day_11/Part_1-10  	   97956	     12201 ns/op	    9608 B/op	     108 allocs/op
Benchmark/Day_11/Part_2-10  	     228	   5249875 ns/op	   10968 B/op	     115 allocs/op
PASS
ok  	github.com/armsnyder/aoc2022	33.485s
```
<!-- END BENCHMARKS -->
