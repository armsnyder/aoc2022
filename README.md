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
Benchmark/Day_01/Part_1-10  	   26493	     45107 ns/op	    4352 B/op	       7 allocs/op
Benchmark/Day_01/Part_2-10  	   26530	     44955 ns/op	    4352 B/op	       7 allocs/op
Benchmark/Day_02/Part_1-10  	   26439	     45608 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_02/Part_2-10  	   28730	     41616 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_03/Part_1-10  	   63450	     17921 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_03/Part_2-10  	   50665	     24905 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_04/Part_1-10  	   19347	     61584 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_04/Part_2-10  	   22399	     53624 ns/op	    4104 B/op	       2 allocs/op
Benchmark/Day_05/Part_1-10  	   16767	     72377 ns/op	   77704 B/op	     538 allocs/op
Benchmark/Day_05/Part_2-10  	   17244	     68840 ns/op	   77728 B/op	     538 allocs/op
PASS
ok  	github.com/armsnyder/aoc2022	16.943s
```
<!-- END BENCHMARKS -->
