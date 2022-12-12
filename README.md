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
Benchmark/Day_01/Part_1-10  	   26564	     45092 ns/op	    4352 B/op	       7 allocs/op
Benchmark/Day_01/Part_2-10  	   26726	     44731 ns/op	    4352 B/op	       7 allocs/op
Benchmark/Day_02/Part_1-10  	   26268	     45336 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_02/Part_2-10  	   28833	     41679 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_03/Part_1-10  	   72262	     16709 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_03/Part_2-10  	   53822	     21414 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_04/Part_1-10  	   19471	     61901 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_04/Part_2-10  	   22471	     53207 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_05/Part_1-10  	   16670	     71521 ns/op	   77704 B/op	     538 allocs/op
Benchmark/Day_05/Part_2-10  	   17428	     68696 ns/op	   77728 B/op	     538 allocs/op
Benchmark/Day_06/Part_1-10  	  155694	      7461 ns/op	     528 B/op	       3 allocs/op
Benchmark/Day_06/Part_2-10  	  104862	     11244 ns/op	     536 B/op	       3 allocs/op
Benchmark/Day_07/Part_1-10  	   20354	     58505 ns/op	   44472 B/op	    1268 allocs/op
Benchmark/Day_07/Part_2-10  	   18616	     64249 ns/op	   48552 B/op	    1277 allocs/op
Benchmark/Day_08/Part_1-10  	   31438	     37828 ns/op	   24584 B/op	       4 allocs/op
Benchmark/Day_08/Part_2-10  	    5756	    190260 ns/op	   14344 B/op	       3 allocs/op
Benchmark/Day_09/Part_1-10  	    1822	    648164 ns/op	  357881 B/op	     191 allocs/op
Benchmark/Day_09/Part_2-10  	    1099	   1080328 ns/op	  177458 B/op	     101 allocs/op
Benchmark/Day_10/Part_1-10  	  365551	      3132 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_10/Part_2-10  	  342951	      3244 ns/op	    4368 B/op	       3 allocs/op
Benchmark/Day_11/Part_1-10  	   96877	     12154 ns/op	    9608 B/op	     108 allocs/op
Benchmark/Day_11/Part_2-10  	     228	   5264237 ns/op	   10968 B/op	     115 allocs/op
Benchmark/Day_12/Part_1-10  	   19963	     58097 ns/op	  125320 B/op	       5 allocs/op
Benchmark/Day_12/Part_2-10  	   25995	     45712 ns/op	  125320 B/op	       5 allocs/op
PASS
ok  	github.com/armsnyder/aoc2022	37.069s
```
<!-- END BENCHMARKS -->
