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
Benchmark/Day_01/Part_1-10  	   26474	     45510 ns/op	    4352 B/op	       7 allocs/op
Benchmark/Day_01/Part_2-10  	   26679	     45106 ns/op	    4352 B/op	       7 allocs/op
Benchmark/Day_02/Part_1-10  	   26244	     45702 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_02/Part_2-10  	   28710	     41887 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_03/Part_1-10  	   87097	     14074 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_03/Part_2-10  	   75066	     14955 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_04/Part_1-10  	   19654	     61127 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_04/Part_2-10  	   22360	     53315 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_05/Part_1-10  	   16533	     71705 ns/op	   77704 B/op	     538 allocs/op
Benchmark/Day_05/Part_2-10  	   17419	     69310 ns/op	   77728 B/op	     538 allocs/op
Benchmark/Day_06/Part_1-10  	  156386	      7481 ns/op	     528 B/op	       3 allocs/op
Benchmark/Day_06/Part_2-10  	  105092	     11289 ns/op	     536 B/op	       3 allocs/op
Benchmark/Day_07/Part_1-10  	   20396	     58760 ns/op	   44472 B/op	    1268 allocs/op
Benchmark/Day_07/Part_2-10  	   18321	     64663 ns/op	   48552 B/op	    1277 allocs/op
Benchmark/Day_08/Part_1-10  	   30986	     38578 ns/op	   24584 B/op	       4 allocs/op
Benchmark/Day_08/Part_2-10  	    5870	    192623 ns/op	   14344 B/op	       3 allocs/op
Benchmark/Day_09/Part_1-10  	    1794	    658102 ns/op	  357877 B/op	     191 allocs/op
Benchmark/Day_09/Part_2-10  	    1094	   1084671 ns/op	  177520 B/op	     101 allocs/op
PASS
ok  	github.com/armsnyder/aoc2022	28.261s
```
<!-- END BENCHMARKS -->
