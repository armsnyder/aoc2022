package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/armsnyder/aoc2022/aocutil"
)

type dayFn func(part2 bool, inputReader io.Reader) any

var days = make(map[int]dayFn)

func declareDay(day int, dayFn dayFn) any {
	days[day] = dayFn
	return nil
}

func main() {
	easternTime, err := time.LoadLocation("EST")
	if err != nil {
		panic(err)
	}

	var (
		day        int
		part2      bool
		skipSubmit bool
		inputFile  string
		benchmark  bool
	)
	flag.IntVar(&day, "d", time.Now().In(easternTime).Add(time.Hour).Day(), "Day of the month")
	flag.BoolVar(&part2, "2", false, "Run part 2")
	flag.BoolVar(&skipSubmit, "s", false, "Skip submitting the answer")
	flag.StringVar(&inputFile, "f", "", "Puzzle input override filepath")
	flag.BoolVar(&benchmark, "b", false, "Run benchmarks and update README.md results")
	flag.Parse()

	if benchmark {
		runBenchmarks()
		return
	}

	dayFn, ok := days[day]
	if !ok {
		fmt.Printf("Generating a stub for day %d\n", day)
		aocutil.GenerateStub(day)
		return
	}

	var input io.ReadCloser
	if inputFile != "" {
		var err error
		input, err = os.Open(inputFile)
		if err != nil {
			panic(err)
		}
	} else {
		input = aocutil.GetInput(day)
	}
	defer input.Close()

	start := time.Now()
	output := dayFn(part2, input)
	runTime := time.Since(start)

	fmt.Printf("Finished in %s\n", runTime)
	fmt.Println("Answer:", output)

	if !skipSubmit {
		aocutil.Submit(day, part2, output)
	}
}

func runBenchmarks() {
	cmd := exec.Command("go", "test", "-benchmem", "-run=^$", "-bench=.")
	cmd.Stderr = os.Stderr
	resultsBuf := &bytes.Buffer{}
	cmd.Stdout = tailWriter{resultsBuf, os.Stdout}
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	original, err := os.ReadFile("README.md")
	if err != nil {
		panic(err)
	}

	start := bytes.Index(original, []byte("<!-- BEGIN BENCHMARKS -->"))
	if start == -1 {
		panic("missing BEGIN BENCHMARKS marker")
	}

	end := bytes.LastIndex(original, []byte("<!-- END BENCHMARKS -->"))
	if end == -1 {
		panic("missing END BENCHMARKS marker")
	}

	fileBuf := &bytes.Buffer{}
	fileBuf.Write(original[:start])
	fileBuf.WriteString("<!-- BEGIN BENCHMARKS -->\n```\n")
	io.Copy(fileBuf, resultsBuf)
	fileBuf.WriteString("```\n")
	fileBuf.Write(original[end:])

	if err := os.WriteFile("README.md", fileBuf.Bytes(), 0o644); err != nil {
		panic(err)
	}
}

type tailWriter struct{ a, b io.Writer }

func (w tailWriter) Write(p []byte) (n int, err error) {
	n, err = w.a.Write(p)
	_, _ = w.b.Write(p)
	return n, err
}
