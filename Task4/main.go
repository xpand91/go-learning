package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	run(os.Stdin, os.Stdout)
}

func run(r io.Reader, w io.Writer) {
	var t int
	fmt.Fscan(r, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(r, &n)

		minTemp, maxTemp := 15, 30
		for j := 0; j < n; j++ {
			var s string
			var a int
			fmt.Fscan(r, &s, &a)

			if strings.HasPrefix(s, ">=") {
				if a > minTemp {
					minTemp = a
				}
			} else if strings.HasPrefix(s, "<=") {
				if a < maxTemp {
					maxTemp = a
				}
			}

			if minTemp <= maxTemp {
				fmt.Fprintln(w, minTemp)
			} else {
				fmt.Fprintln(w, -1)
			}
		}
		fmt.Fprintln(w)
	}
}
