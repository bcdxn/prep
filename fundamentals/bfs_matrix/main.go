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
	r int
	c int
}

func shortestPathBFS(m [][]int) int {
	q := queue.New[rc]()
	visited := make(map[rc]struct{})

	rows := len(m)
	if rows < 1 {
		return -1
	}

	cols := len(m[0])
	if cols < 1 {
		return -1
	}

	q.Enqueue(rc{r: 0, c: 0})
	visited[rc{r: 0, c: 0}] = struct{}{}

	length := 0

	for q.Size() > 0 {
		for range q.Size() {
			node := *q.Dequeue()

			if node.r == rows-1 && node.c == cols-1 {
				return length
			}

			// add neighbors to the queue
			if _, ok := visited[rc{r: node.r + 1, c: node.c}]; !ok && node.r+1 < rows && m[node.r+1][node.c] != 1 {
				q.Enqueue(rc{r: node.r + 1, c: node.c})
				visited[rc{r: node.r + 1, c: node.c}] = struct{}{}
			}
			if _, ok := visited[rc{r: node.r - 1, c: node.c}]; !ok && node.r-1 >= 0 && m[node.r-1][node.c] != 1 {
				q.Enqueue(rc{r: node.r - 1, c: node.c})
				visited[rc{r: node.r - 1, c: node.c}] = struct{}{}
			}
			if _, ok := visited[rc{r: node.r, c: node.c + 1}]; !ok && node.c+1 < cols && m[node.r][node.c+1] != 1 {
				q.Enqueue(rc{r: node.r, c: node.c + 1})
				visited[rc{r: node.r, c: node.c + 1}] = struct{}{}
			}
			if _, ok := visited[rc{r: node.r, c: node.c - 1}]; !ok && node.c-1 >= cols && m[node.r][node.c-1] != 1 {
				q.Enqueue(rc{r: node.r, c: node.c - 1})
				visited[rc{r: node.r, c: node.c - 1}] = struct{}{}
			}
		}
		length += 1
	}

	return -1
}
