package main

import (
	"errors"
	"fmt"
)

func main() {
	in := [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 1},
		{0, 0, 1, 0, 1, 0},
		{1, 1, 0, 0, 1, 0},
		{1, 0, 1, 1, 0, 0},
		{1, 0, 0, 0, 0, 1},
	}

	// out := [][]int{
	// 	{1, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 1, 1, 1},
	// 	{0, 0, 0, 0, 1, 0},
	// 	{1, 1, 0, 0, 1, 0},
	// 	{1, 0, 0, 0, 0, 0},
	// 	{1, 0, 0, 0, 0, 1},
	// }

	flipIslands(in)

	for i := range len(in) {
		row := make([]int, 0)
		for j := range len(in[i]) {
			row = append(row, in[i][j])
		}
		fmt.Println(row)
	}
}

// record all paths that start on the board and end not on the boarder

func flipIslands(m [][]int) {
	rows := len(m)
	if rows < 3 {
		return
	}

	cols := len(m[0])
	if cols < 3 {
		return
	}

	visited := make(map[rc]struct{})

	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			island := make(map[rc]struct{})
			if _, ok := visited[rc{r, c}]; !ok && m[r][c] == 1 {
				// begin test
				q := queue{}
				q.Enqueue(rc{r, c})
				island[rc{r, c}] = struct{}{}
				isIsland := true
				// bfs from current r,c
				for q.Size() > 0 {
					for range q.Size() {
						n, _ := q.Dequeue()

						if _, ok := island[rc{n.r - 1, n.c}]; !ok && n.r-1 >= 0 && m[n.r-1][n.c] == 1 {
							q.Enqueue(rc{n.r - 1, n.c})
							island[rc{n.r - 1, n.c}] = struct{}{}
							visited[rc{n.r - 1, n.c}] = struct{}{}
							if n.r-1 == 0 {
								isIsland = false
							}
						}

						if _, ok := island[rc{n.r + 1, n.c}]; !ok && n.r+1 < rows && m[n.r+1][n.c] == 1 {
							q.Enqueue(rc{n.r + 1, n.c})
							island[rc{n.r + 1, n.c}] = struct{}{}
							visited[rc{n.r + 1, n.c}] = struct{}{}
							if n.r+1 == rows-1 {
								isIsland = false
							}
						}

						if _, ok := island[rc{n.r, n.c - 1}]; !ok && n.c-1 >= 0 && m[n.r][n.c-1] == 1 {
							q.Enqueue(rc{n.r, n.c - 1})
							island[rc{n.r, n.c - 1}] = struct{}{}
							visited[rc{n.r, n.c - 1}] = struct{}{}
							if n.c-1 == 0 {
								isIsland = false
							}
						}

						if _, ok := island[rc{n.r, n.c + 1}]; !ok && n.c+1 < cols && m[n.r][n.c+1] == 1 {
							q.Enqueue(rc{n.r, n.c + 1})
							island[rc{n.r, n.c + 1}] = struct{}{}
							visited[rc{n.r, n.c + 1}] = struct{}{}
							if n.c+1 == cols-1 {
								isIsland = false
							}
						}
					}

				}

				if isIsland {
					// flip the bits in the islands
					for n := range island {
						m[n.r][n.c] = 0
					}
				}
			}
		}
	}
}

type rc struct {
	r int
	c int
}

type queue struct {
	s []rc
}

func (q *queue) Enqueue(item rc) {
	q.s = append(q.s, item)
}

func (q *queue) Dequeue() (rc, error) {
	var node rc
	if q.Size() < 1 {
		return node, errors.New("cannot perform Dequeue on empty queue")
	}

	node = q.s[0]
	q.s = q.s[1:]

	return node, nil
}

func (q queue) Size() int {
	return len(q.s)
}
