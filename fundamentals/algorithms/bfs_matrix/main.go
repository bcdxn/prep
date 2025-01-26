package main

import (
	"errors"
	"fmt"
)

func main() {
	matrix := [][]int{
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 1},
		{0, 1, 0, 0},
	}

	shortestPath := bfs(matrix)
	fmt.Println("shortest path: ", shortestPath)
}

func bfs(m [][]int) int {
	rows := len(m)
	cols := len(m[0])

	q := queue{}
	visited := make(map[rc]struct{})

	q.Enqueue(rc{0, 0})
	visited[rc{0, 0}] = struct{}{}
	length := 0

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

type rc struct {
	r int
	c int
}

type queue struct {
	q []rc
}

func (q *queue) Enqueue(node rc) {
	q.q = append(q.q, node)
}

func (q *queue) Dequeue() (rc, error) {
	if len(q.q) < 1 {
		return rc{}, errors.New("cannot perform dequeue on empty queue")
	}

	node := q.q[0]
	q.q = q.q[1:]

	return node, nil
}

func (q queue) Size() int {
	return len(q.q)
}
