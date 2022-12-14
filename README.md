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
Benchmark/Day_01/Part_1-10  	   27812	     42900 ns/op	    4352 B/op	       7 allocs/op
Benchmark/Day_01/Part_2-10  	   27996	     42707 ns/op	    4352 B/op	       7 allocs/op
Benchmark/Day_02/Part_1-10  	   27171	     43970 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_02/Part_2-10  	   30682	     38885 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_03/Part_1-10  	   65037	     16574 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_03/Part_2-10  	   57139	     18163 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_04/Part_1-10  	   19728	     60718 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_04/Part_2-10  	   22538	     53213 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_05/Part_1-10  	   16809	     71400 ns/op	   77704 B/op	     538 allocs/op
Benchmark/Day_05/Part_2-10  	   17222	     68599 ns/op	   77728 B/op	     538 allocs/op
Benchmark/Day_06/Part_1-10  	  153066	      7507 ns/op	     528 B/op	       3 allocs/op
Benchmark/Day_06/Part_2-10  	  105386	     11215 ns/op	     536 B/op	       3 allocs/op
Benchmark/Day_07/Part_1-10  	   20755	     57663 ns/op	   44472 B/op	    1268 allocs/op
Benchmark/Day_07/Part_2-10  	   18920	     63186 ns/op	   48552 B/op	    1277 allocs/op
Benchmark/Day_08/Part_1-10  	   31466	     37536 ns/op	   24584 B/op	       4 allocs/op
Benchmark/Day_08/Part_2-10  	    5821	    194676 ns/op	   14344 B/op	       3 allocs/op
Benchmark/Day_09/Part_1-10  	    1768	    647177 ns/op	  357909 B/op	     191 allocs/op
Benchmark/Day_09/Part_2-10  	    1095	   1077486 ns/op	  177583 B/op	     101 allocs/op
Benchmark/Day_10/Part_1-10  	  376341	      3046 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_10/Part_2-10  	  362672	      3185 ns/op	    4368 B/op	       3 allocs/op
Benchmark/Day_11/Part_1-10  	   95172	     12488 ns/op	    9608 B/op	     108 allocs/op
Benchmark/Day_11/Part_2-10  	     225	   5294519 ns/op	   10968 B/op	     115 allocs/op
Benchmark/Day_12/Part_1-10  	   17494	     68208 ns/op	  125320 B/op	       5 allocs/op
Benchmark/Day_12/Part_2-10  	   26238	     45275 ns/op	  125320 B/op	       5 allocs/op
Benchmark/Day_13/Part_1-10  	    2605	    444877 ns/op	  716970 B/op	    8676 allocs/op
Benchmark/Day_13/Part_2-10  	    2202	    514286 ns/op	  717099 B/op	    8680 allocs/op
PASS
ok  	github.com/armsnyder/aoc2022	39.219s
```
<!-- END BENCHMARKS -->
