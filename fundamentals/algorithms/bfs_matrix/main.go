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

	length := bfs(m)

	fmt.Println("shortest path: ", length) // should print 6
}

type t struct {
	r int
	c int
}

func bfs(m [][]int) int {
	rows := len(m)
	if rows < 1 {
		return -1
	}

	cols := len(m[0])
	if cols < 1 {
		return -1
	}

	q := queue.New[t]()
	visited := make(map[t]struct{})

	q.Enqueue(t{0, 0})
	visited[t{0, 0}] = struct{}{}

	length := 0

	for q.Size() > 0 {
		for range q.Size() {
			n := q.Dequeue()

			if n.r == rows-1 && n.c == cols-1 {
				return length
			}

			if _, ok := visited[t{n.r - 1, n.c}]; !ok && n.r-1 >= 0 && m[n.r-1][n.c] != 1 {
				q.Enqueue(t{n.r - 1, n.c})
				visited[t{n.r - 1, n.c}] = struct{}{}
			}
			if _, ok := visited[t{n.r + 1, n.c}]; !ok && n.r+1 < rows && m[n.r+1][n.c] != 1 {
				q.Enqueue(t{n.r + 1, n.c})
				visited[t{n.r + 1, n.c}] = struct{}{}
			}
			if _, ok := visited[t{n.r, n.c - 1}]; !ok && n.c-1 >= 0 && m[n.r][n.c-1] != 1 {
				q.Enqueue(t{n.r, n.c - 1})
				visited[t{n.r, n.c - 1}] = struct{}{}
			}
			if _, ok := visited[t{n.r, n.c + 1}]; !ok && n.c+1 < cols && m[n.r][n.c+1] != 1 {
				q.Enqueue(t{n.r, n.c + 1})
				visited[t{n.r, n.c + 1}] = struct{}{}
			}

		}
		length += 1
	}

	return -1
}
