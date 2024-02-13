package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	run(os.Stdin, os.Stdout)
}

func run(r io.Reader, w io.Writer) {
	var t int
	fmt.Fscan(r, &t)

	for i := 0; i < t; i++ {
		var d, m, y int
		fmt.Fscan(r, &d, &m, &y)
		fmt.Fprintln(w, isValidDate(d, m, y))
	}
}

func isValidDate(d int, m int, y int) string {
	t := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
	if t.Day() != d || t.Month() != time.Month(m) || t.Year() != y {
		return "NO"
	}
	return "YES"
}
