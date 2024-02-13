package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	run(os.Stdin, os.Stdout)
}

func run(r io.Reader, w io.Writer) {
	var t int
	fmt.Fscan(r, &t)

	for i := 0; i < t; i++ {
		var s string
		fmt.Fscan(r, &s)
		fmt.Fprintln(w, splitCarNumbers(s))
	}
}

func splitCarNumbers(s string) string {
	r1 := regexp.MustCompile(`^[A-Z][0-9]{2}[A-Z]{2}`)
	r2 := regexp.MustCompile(`^[A-Z][0-9][A-Z]{2}`)

	var result []string
	for len(s) > 0 {
		if r1.MatchString(s) {
			result = append(result, s[:5])
			s = s[5:]
		} else if r2.MatchString(s) {
			result = append(result, s[:4])
			s = s[4:]
		} else {
			return "-"
		}
	}

	return strings.Join(result, " ")
}
