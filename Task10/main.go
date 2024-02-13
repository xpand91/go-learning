package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Comment struct {
	id   int
	text string
}

func main() {
	run(os.Stdin, os.Stdout)
}

func run(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	writer := bufio.NewWriter(w)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		comments := make(map[int]Comment)
		children := make(map[int][]int)

		for j := 0; j < n; j++ {
			scanner.Scan()
			line := scanner.Text()
			parts := strings.SplitN(line, " ", 3)
			id, _ := strconv.Atoi(parts[0])
			parent, _ := strconv.Atoi(parts[1])
			text := parts[2]

			comments[id] = Comment{id, text}
			children[parent] = append(children[parent], id)
		}

		for _, v := range children {
			sort.Ints(v)
		}

		printComments(writer, children, comments, -1, "")
	}

	writer.Flush()
}

func printComments(w io.Writer, children map[int][]int, comments map[int]Comment, id int, prefix string) {
	for i, child := range children[id] {
		isLast := i == len(children[id])-1
		if id == -1 {
			fmt.Fprintf(w, "\n%s%s\n", prefix, comments[child].text)
		} else {
			fmt.Fprintf(w, "%s|\n%s|--%s\n", prefix[3:], prefix[3:], comments[child].text)
		}

		if isLast {
			printComments(w, children, comments, child, prefix+"   ")
		} else {
			printComments(w, children, comments, child, prefix+"|  ")
		}
	}
}
