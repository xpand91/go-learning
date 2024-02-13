package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Terminal struct {
	lines           []string
	curLine, curPos int
}

func NewTerminal() *Terminal {
	return &Terminal{lines: []string{""}}
}

func (t *Terminal) Insert(ch byte) {
	line := t.lines[t.curLine]
	t.lines[t.curLine] = line[:t.curPos] + string(ch) + line[t.curPos:]
	t.curPos++
}

func (t *Terminal) Left() {
	if t.curPos > 0 {
		t.curPos--
	}
}

func (t *Terminal) Right() {
	if t.curPos < len(t.lines[t.curLine]) {
		t.curPos++
	}
}

func (t *Terminal) Up() {
	if t.curLine > 0 {
		t.curLine--
		if t.curPos > len(t.lines[t.curLine]) {
			t.curPos = len(t.lines[t.curLine])
		}
	}
}

func (t *Terminal) Down() {
	if t.curLine < len(t.lines)-1 {
		t.curLine++
		if t.curPos > len(t.lines[t.curLine]) {
			t.curPos = len(t.lines[t.curLine])
		}
	}
}

func (t *Terminal) Home() {
	t.curPos = 0
}

func (t *Terminal) End() {
	t.curPos = len(t.lines[t.curLine])
}

func (t *Terminal) Enter() {
	line := t.lines[t.curLine]
	t.lines[t.curLine] = line[:t.curPos]
	t.lines = append(t.lines[:t.curLine+1], append([]string{line[t.curPos:]}, t.lines[t.curLine+1:]...)...)
	t.curLine++
	t.curPos = 0
}

func (t *Terminal) Process(ch byte) {
	switch ch {
	case 'L':
		t.Left()
	case 'R':
		t.Right()
	case 'U':
		t.Up()
	case 'D':
		t.Down()
	case 'B':
		t.Home()
	case 'E':
		t.End()
	case 'N':
		t.Enter()
	default:
		t.Insert(ch)
	}
}

func (t *Terminal) Print(w io.Writer) {
	for _, line := range t.lines {
		fmt.Fprintln(w, line)
	}
	fmt.Fprintln(w, "-")
}

func run(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		scanner.Scan()
		term := NewTerminal()
		for _, ch := range scanner.Text() {
			term.Process(byte(ch))
		}
		term.Print(w)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
