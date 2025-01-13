package main

import (
	"fmt"

	ds "github.com/bcdxn/leetcode/fundamentals/data_structures"
)

var matrix = [][]int{
	{0, 0, 0, 0},
	{1, 1, 0, 0},
	{0, 0, 0, 1},
	{0, 1, 0, 0},
}

func main() {
	shortestPath := bfs(matrix)
	fmt.Println("shortest path: ", shortestPath)
}

type rc struct {
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

	length := 0

	q := ds.NewQueue[rc]()
	visited := make(map[rc]struct{})

	q.Enqueue(rc{0, 0})
	visited[rc{0, 0}] = struct{}{}

	for q.Size() > 0 {
		for range q.Size() {
			n, _ := q.Dequeue()

			if n.r == rows-1 && n.c == cols-1 {
				return length
			}

			if _, ok := visited[rc{n.r - 1, n.c}]; !ok && n.r-1 >= 0 && m[n.r-1][n.c] != 1 {
				q.Enqueue(rc{n.r - 1, n.c})
				visited[rc{n.r - 1, n.c}] = struct{}{}
			}
			if _, ok := visited[rc{n.r + 1, n.c}]; !ok && n.r+1 < rows && m[n.r+1][n.c] != 1 {
				q.Enqueue(rc{n.r + 1, n.c})
				visited[rc{n.r + 1, n.c}] = struct{}{}
			}
			if _, ok := visited[rc{n.r, n.c - 1}]; !ok && n.c-1 >= 0 && m[n.r][n.c-1] != 1 {
				q.Enqueue(rc{n.r, n.c - 1})
				visited[rc{n.r, n.c - 1}] = struct{}{}
			}
			if _, ok := visited[rc{n.r, n.c + 1}]; !ok && n.c+1 < cols && m[n.r][n.c+1] != 1 {
				q.Enqueue(rc{n.r, n.c + 1})
				visited[rc{n.r, n.c + 1}] = struct{}{}
			}
		}
		length += 1
	}

	return -1
}
