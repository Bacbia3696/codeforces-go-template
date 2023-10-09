package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func scan(a ...any) {
	fmt.Fscan(reader, a...)
}

func println(a ...any) {
	fmt.Fprintln(writer, a...)
}

func printf(format string, a ...any) {
	fmt.Fprintf(writer, format, a...)
}

func max(x int, y ...int) int {
	for _, v := range y {
		if v > x {
			x = v
		}
	}
	return x
}

func min(x int, y ...int) int {
	for _, v := range y {
		if v < x {
			x = v
		}
	}
	return x
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	t := 1

	if MULTIPLE_TEST_CASES {
		scan(&t)
	}
	for ; t > 0; t-- {
		solve()
	}
	writer.Flush()
}

const MULTIPLE_TEST_CASES = true

func solve() {

}
