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
Benchmark/Day_01/Part_1-10  	   26460	     45251 ns/op	    4352 B/op	       7 allocs/op
Benchmark/Day_01/Part_2-10  	   26678	     44943 ns/op	    4352 B/op	       7 allocs/op
Benchmark/Day_02/Part_1-10  	   26300	     45621 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_02/Part_2-10  	   28716	     41779 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_03/Part_1-10  	   85351	     13909 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_03/Part_2-10  	   76328	     14960 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_04/Part_1-10  	   19537	     62451 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_04/Part_2-10  	   22063	     54152 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_05/Part_1-10  	   16710	     72308 ns/op	   77704 B/op	     538 allocs/op
Benchmark/Day_05/Part_2-10  	   17269	     69682 ns/op	   77728 B/op	     538 allocs/op
Benchmark/Day_06/Part_1-10  	  155301	      7464 ns/op	     528 B/op	       3 allocs/op
Benchmark/Day_06/Part_2-10  	  105946	     11338 ns/op	     536 B/op	       3 allocs/op
Benchmark/Day_07/Part_1-10  	   20326	     58280 ns/op	   44472 B/op	    1268 allocs/op
Benchmark/Day_07/Part_2-10  	   18510	     64688 ns/op	   48552 B/op	    1277 allocs/op
Benchmark/Day_08/Part_1-10  	   31150	     39285 ns/op	   24584 B/op	       4 allocs/op
Benchmark/Day_08/Part_2-10  	    5737	    192811 ns/op	   14344 B/op	       3 allocs/op
PASS
ok  	github.com/armsnyder/aoc2022	25.725s
```
<!-- END BENCHMARKS -->
