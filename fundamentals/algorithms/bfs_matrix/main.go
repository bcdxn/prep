package main

import (
	"fmt"
	"strings"

	"github.com/bcdxn/leetcode/fundamentals/algorithms/queue"
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

/* Simple queue implementation
------------------------------------------------------------------------------------------------- */

type Queue[T any] struct {
	s []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(item T) {
	q.s = append(q.s, item)
}

func (q *Queue[T]) Peek() *T {
	if len(q.s) < 1 {
		return nil
	}

	return &q.s[0]
}

func (q *Queue[T]) Dequeue() *T {
	if len(q.s) < 1 {
		return nil
	}
	// store first item in queue to return
	item := q.s[0]
	// remove first item from queue
	q.s = q.s[1:]
	// return stored item
	return &item
}

func (q *Queue[T]) String() string {
	var str []string
	for _, val := range q.s {
		str = append(str, fmt.Sprintf("%v", val))
	}

	return "[" + strings.Join(str, " ") + "]"
}

func (q *Queue[T]) Size() int {
	return len(q.s)
}
