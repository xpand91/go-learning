package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	run(os.Stdin, os.Stdout)
}

func run(r io.Reader, w io.Writer) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	wr := bufio.NewWriter(w)
	sc.Scan()
	t, _ := strconv.Atoi(sc.Text())
	for i := 0; i < t; i++ {
		sc.Scan()
		n, _ := strconv.Atoi(sc.Text())
		a := make([]int, n)
		for j := 0; j < n; j++ {
			sc.Scan()
			a[j], _ = strconv.Atoi(sc.Text())
		}
		res := compress(a)
		fmt.Fprintln(wr, len(res))
		for i, v := range res {
			if i != 0 {
				fmt.Fprint(wr, " ")
			}
			fmt.Fprint(wr, strconv.Itoa(v))
		}
		fmt.Fprintln(wr)
	}
	wr.Flush()
}

func compress(a []int) []int {
	res := []int{a[0]}
	count := 0
	inc := true
	if len(a) > 1 {
		inc = a[1] >= a[0]
	}
	for i := 1; i < len(a); i++ {
		if (a[i] == (a[i-1]+1) && inc) || (a[i] == (a[i-1]-1) && !inc) {
			count++
		} else {
			if inc {
				res = append(res, count)
			} else {
				res = append(res, -count)
			}
			res = append(res, a[i])
			count = 0
			if i+1 < len(a) {
				inc = a[i+1] >= a[i]
			}
		}
	}
	if inc {
		res = append(res, count)
	} else {
		res = append(res, -count)
	}
	return res
}
