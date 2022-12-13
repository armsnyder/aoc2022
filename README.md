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
Benchmark/Day_01/Part_1-10  	   26354	     45319 ns/op	    4352 B/op	       7 allocs/op
Benchmark/Day_01/Part_2-10  	   26653	     44845 ns/op	    4352 B/op	       7 allocs/op
Benchmark/Day_02/Part_1-10  	   26209	     45429 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_02/Part_2-10  	   28836	     41772 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_03/Part_1-10  	   72885	     16679 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_03/Part_2-10  	   55695	     20019 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_04/Part_1-10  	   19610	     61124 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_04/Part_2-10  	   22556	     53143 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_05/Part_1-10  	   16742	     71565 ns/op	   77704 B/op	     538 allocs/op
Benchmark/Day_05/Part_2-10  	   17430	     68766 ns/op	   77728 B/op	     538 allocs/op
Benchmark/Day_06/Part_1-10  	  153930	      7521 ns/op	     528 B/op	       3 allocs/op
Benchmark/Day_06/Part_2-10  	  105321	     11257 ns/op	     536 B/op	       3 allocs/op
Benchmark/Day_07/Part_1-10  	   20407	     58779 ns/op	   44472 B/op	    1268 allocs/op
Benchmark/Day_07/Part_2-10  	   18478	     64801 ns/op	   48552 B/op	    1277 allocs/op
Benchmark/Day_08/Part_1-10  	   31687	     37452 ns/op	   24584 B/op	       4 allocs/op
Benchmark/Day_08/Part_2-10  	    5780	    190597 ns/op	   14344 B/op	       3 allocs/op
Benchmark/Day_09/Part_1-10  	    1792	    647840 ns/op	  357843 B/op	     191 allocs/op
Benchmark/Day_09/Part_2-10  	    1102	   1078948 ns/op	  177496 B/op	     101 allocs/op
Benchmark/Day_10/Part_1-10  	  370096	      3145 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_10/Part_2-10  	  348694	      3286 ns/op	    4368 B/op	       3 allocs/op
Benchmark/Day_11/Part_1-10  	   93002	     12269 ns/op	    9608 B/op	     108 allocs/op
Benchmark/Day_11/Part_2-10  	     226	   5355633 ns/op	   10968 B/op	     115 allocs/op
Benchmark/Day_12/Part_1-10  	   19596	     57800 ns/op	  125320 B/op	       5 allocs/op
Benchmark/Day_12/Part_2-10  	   26005	     45731 ns/op	  125320 B/op	       5 allocs/op
Benchmark/Day_13/Part_1-10  	   67928	     16496 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_13/Part_2-10  	    1380	    856051 ns/op	   27640 B/op	     304 allocs/op
PASS
ok  	github.com/armsnyder/aoc2022	40.638s
```
<!-- END BENCHMARKS -->
