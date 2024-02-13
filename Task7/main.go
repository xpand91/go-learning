package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	run(os.Stdin, os.Stdout)
}

func run(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	writer := bufio.NewWriter(w)
	defer writer.Flush()

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		printedPages := scanner.Text()

		pages := make([]bool, k+1)
		for _, pageRange := range strings.Split(printedPages, ",") {
			if strings.Contains(pageRange, "-") {
				bounds := strings.Split(pageRange, "-")
				l, _ := strconv.Atoi(bounds[0])
				r, _ := strconv.Atoi(bounds[1])
				for j := l; j <= r; j++ {
					pages[j] = true
				}
			} else {
				page, _ := strconv.Atoi(pageRange)
				pages[page] = true
			}
		}

		start := -1
		var isFirst = true
		for j := 1; j <= k; j++ {
			if !pages[j] {
				if start == -1 {
					start = j
				}
				if j == k || pages[j+1] {
					if isFirst {
						isFirst = false
					} else {
						fmt.Fprint(writer, ",")
					}
					if start == j {
						fmt.Fprintf(writer, "%d", start)
					} else {
						fmt.Fprintf(writer, "%d-%d", start, j)
					}
					start = -1
				}
			}
		}
		writer.WriteByte('\n')
	}
}
