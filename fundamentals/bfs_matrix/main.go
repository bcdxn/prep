package main

import (
	"fmt"

	"github.com/bcdxn/leetcode/fundamentals/queue"
)

func main() {
	m := [][]int{
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 1},
		{0, 1, 0, 0},
	}

	length := shortestPathBFS(m)

	fmt.Println("shortest path: ", length) // should print 6
}

type rc struct {
	R int
	C int
}

func shortestPathBFS(m [][]int) int {
	rows := len(m)
	if rows < 1 {
		return 0
	}
	cols := len(m[0])
	if cols < 1 {
		return 0
	}

	q := queue.New[rc]()
	visited := make(map[rc]struct{})
	q.Enqueue(rc{R: 0, C: 0})
	visited[rc{R: 0, C: 0}] = struct{}{}

	length := 0

	for q.Size() > 0 {
		size := q.Size()
		for i := 0; i < size; i++ {
			t := *q.Dequeue()
			r, c := t.R, t.C
			// reached the destination base case
			if r == rows-1 && c == rows-1 {
				return length
			}
			// visit neighbors (checking bounds, checking its not blocked, and checking its not already visited)
			if _, ok := visited[rc{R: r - 1, C: c}]; !ok && r-1 >= 0 && m[r-1][c] != 1 {
				q.Enqueue(rc{R: r - 1, C: c})
				visited[rc{R: r - 1, C: c}] = struct{}{}
			}
			if _, ok := visited[rc{R: r + 1, C: c}]; !ok && r+1 < rows && m[r+1][c] != 1 {
				q.Enqueue(rc{R: r + 1, C: c})
				visited[rc{R: r + 1, C: c}] = struct{}{}
			}

			if _, ok := visited[rc{R: r, C: c - 1}]; !ok && c-1 >= 0 && m[r][c-1] != 1 {
				q.Enqueue(rc{R: r, C: c - 1})
				visited[rc{R: r, C: c - 1}] = struct{}{}
			}
			if _, ok := visited[rc{R: r, C: c + 1}]; !ok && c+1 < cols && m[r][c+1] != 1 {
				q.Enqueue(rc{R: r, C: c + 1})
				visited[rc{R: r, C: c + 1}] = struct{}{}
			}
		}
		length += 1
	}

	return length
}
