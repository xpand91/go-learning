package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type pair struct{ x, y int }

var (
	dx = []int{-1, 0, 1, 0}
	dy = []int{0, 1, 0, -1}
)

func main() {
	run(os.Stdin, os.Stdout)
}

func run(r io.Reader, w io.Writer) {
	var t int
	fmt.Fscan(r, &t)
	for t > 0 {
		solve(r, w)
		t--
	}
}

func solve(r io.Reader, w io.Writer) {
	var n, m int
	fmt.Fscan(r, &n, &m)
	grid := make([][]byte, n)
	for i := range grid {
		fmt.Fscan(r, &grid[i])
	}
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}
	frames := [][]pair{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '*' && !visited[i][j] {
				frames = append(frames, bfs(i, j, grid, visited))
			}
		}
	}
	depth := make([]int, len(frames))
	for i, frame := range frames {
		for j, other := range frames {
			if i != j && isContained(frame, other) {
				depth[i] = max(depth[i], depth[j]+1)
			}
		}
	}
	sort.Ints(depth)
	for _, v := range depth {
		fmt.Fprint(w, v, " ")
	}
	fmt.Fprintln(w)
}

func bfs(x, y int, grid [][]byte, visited [][]bool) []pair {
	n := len(grid)
	m := len(grid[0])
	queue := []pair{{x, y}}
	frame := []pair{{x, y}}
	visited[x][y] = true
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			nx, ny := cell.x+dx[i], cell.y+dy[i]
			if isInside(nx, ny, n, m) && grid[nx][ny] == '*' && !visited[nx][ny] {
				visited[nx][ny] = true
				queue = append(queue, pair{nx, ny})
				frame = append(frame, pair{nx, ny})
			}
		}
	}
	return frame
}

func isInside(x, y, n, m int) bool {
	return x >= 0 && y >= 0 && x < n && y < m
}

func isContained(frame1, frame2 []pair) bool {
	minX, minY, maxX, maxY := frame2[0].x, frame2[0].y, frame2[0].x, frame2[0].y
	for _, c := range frame2 {
		minX = min(minX, c.x)
		minY = min(minY, c.y)
		maxX = max(maxX, c.x)
		maxY = max(maxY, c.y)
	}
	for _, c1 := range frame1 {
		if c1.x < minX || c1.y < minY || c1.x > maxX || c1.y > maxY {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
