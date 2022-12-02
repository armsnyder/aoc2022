package aocutil

import (
	"bufio"
	"io"
	"strconv"
	"unicode/utf8"
)

func ReadAllInts(inputReader io.Reader) []int {
	var result []int
	VisitInts(inputReader, func(v int) {
		result = append(result, v)
	})
	return result
}

func VisitInts(inputReader io.Reader, visitFn func(v int)) {
	VisitLines(inputReader, func(v []byte) {
		if len(v) > 0 {
			n, err := AtoiBytes(v)
			if err != nil {
				panic(err)
			}
			visitFn(n)
		}
	})
}

func ReadAllStrings(inputReader io.Reader) []string {
	var result []string
	VisitStrings(inputReader, func(v []byte) {
		result = append(result, string(v))
	})
	return result
}

func Read2DByteArray(inputReader io.Reader) [][]byte {
	var result [][]byte
	scanner := bufio.NewScanner(inputReader)
	for scanner.Scan() {
		if len(scanner.Bytes()) > 0 {
			row := make([]byte, len(scanner.Bytes()))
			copy(row, scanner.Bytes())
			result = append(result, row)
		}
	}
	return result
}

func ReadAllCommaSeparatedInts(inputReader io.Reader) []int {
	var result []int
	VisitCommaSeparatedInts(inputReader, func(n int) {
		result = append(result, n)
	})
	return result
}

func VisitCommaSeparatedInts(inputReader io.Reader, visitFn func(v int)) {
	VisitCommaSeparated(inputReader, func(v []byte) {
		n, err := AtoiBytes(v)
		if err != nil {
			panic(err)
		}
		visitFn(n)
	})
}

func VisitStrings(inputReader io.Reader, visitFn func(v []byte)) {
	VisitLines(inputReader, func(v []byte) {
		if len(v) > 0 {
			visitFn(v)
		}
	})
}

func VisitStringGroups(inputReader io.Reader, visitFn func(v []string)) {
	var group []string
	VisitLines(inputReader, func(v []byte) {
		if len(v) > 0 {
			group = append(group, string(v))
		} else if len(group) > 0 {
			visitFn(group)
			group = group[:0]
		}
	})
	if len(group) > 0 {
		visitFn(group)
	}
}

func VisitIntGroups(inputReader io.Reader, visitFn func(v []int)) {
	var group []int
	VisitLines(inputReader, func(v []byte) {
		if len(v) > 0 {
			n, err := AtoiBytes(v)
			if err != nil {
				panic(err)
			}
			group = append(group, n)
		} else if len(group) > 0 {
			visitFn(group)
			group = group[:0]
		}
	})
	if len(group) > 0 {
		visitFn(group)
	}
}

func VisitLines(inputReader io.Reader, visitFn func(v []byte)) {
	scanner := bufio.NewScanner(inputReader)
	for scanner.Scan() {
		visitFn(scanner.Bytes())
	}
}

func VisitCommaSeparated(inputReader io.Reader, visitFn func(v []byte)) {
	scanner := bufio.NewScanner(inputReader)
	scanner.Split(ScanCommaSeparated)
	for scanner.Scan() {
		visitFn(scanner.Bytes())
	}
}

// ScanCommaSeparated is a split function for a Scanner that returns each comma-separated word of
// text.
func ScanCommaSeparated(data []byte, atEOF bool) (advance int, token []byte, err error) {
	isDelimiter := func(r rune) bool {
		switch r {
		case ' ', '\n', ',':
			return true
		}
		return false
	}
	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !isDelimiter(r) {
			break
		}
	}
	// Scan until space, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if isDelimiter(r) {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

// AtoiBytes is based on strconv.Atoi() but skips the need to convert to string, thereby reducing
// memory allocations.
func AtoiBytes(s []byte) (int, error) {
	const fnAtoiBytes = "AtoiBytes"

	sLen := len(s)
	if strconv.IntSize == 32 && (0 < sLen && sLen < 10) ||
		strconv.IntSize == 64 && (0 < sLen && sLen < 19) {
		// Fast path for small integers that fit int type.
		s0 := s
		if s[0] == '-' || s[0] == '+' {
			s = s[1:]
			if len(s) < 1 {
				return 0, &strconv.NumError{Func: fnAtoiBytes, Num: string(s0), Err: strconv.ErrSyntax}
			}
		}

		n := 0
		for _, ch := range []byte(s) {
			ch -= '0'
			if ch > 9 {
				return 0, &strconv.NumError{Func: fnAtoiBytes, Num: string(s0), Err: strconv.ErrSyntax}
			}
			n = n*10 + int(ch)
		}
		if s0[0] == '-' {
			n = -n
		}
		return n, nil
	}

	// Slow path for invalid, big, or underscored integers.
	i64, err := strconv.ParseInt(string(s), 10, 0)
	if nerr, ok := err.(*strconv.NumError); ok {
		nerr.Func = fnAtoiBytes
	}
	return int(i64), err
}
