package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	run(os.Stdin, os.Stdout)
}

func run(r io.Reader, w io.Writer) {
	var t int
	fmt.Fscan(r, &t)
	for i := 0; i < t; i++ {
		ships := make([]int, 10)
		for j := 0; j < 10; j++ {
			fmt.Fscan(r, &ships[j])
		}
		sort.Ints(ships)
		if ships[0] == 1 && ships[1] == 1 && ships[2] == 1 && ships[3] == 1 &&
			ships[4] == 2 && ships[5] == 2 && ships[6] == 2 &&
			ships[7] == 3 && ships[8] == 3 &&
			ships[9] == 4 {
			fmt.Fprintln(w, "YES")
		} else {
			fmt.Fprintln(w, "NO")
		}
	}
}
