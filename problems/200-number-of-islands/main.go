package main

import "fmt"

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	count := numIslands(grid)
	fmt.Println("count: ", count) // 1
}

func numIslands(grid [][]byte) int {
	total := 0
	visited := make(map[rc]struct{})

	for r := range grid {
		for c := range grid[r] {
			if dfs(grid, rc{r, c}, visited) {
				total++
			}
		}
	}

	return total
}

func dfs(m [][]byte, coord rc, visited map[rc]struct{}) bool {
	rows := len(m)
	cols := len(m[0])
	// out of bounds
	if coord.r < 0 || coord.r >= rows || coord.c < 0 || coord.c >= cols {
		return false
	}
	// ocean
	if m[coord.r][coord.c] == '0' {
		return false
	}
	// already visited
	if _, ok := visited[coord]; ok {
		return false
	}

	// mark as visited and recurse
	visited[coord] = struct{}{}

	dfs(m, rc{coord.r - 1, coord.c}, visited)
	dfs(m, rc{coord.r, coord.c + 1}, visited)
	dfs(m, rc{coord.r + 1, coord.c}, visited)
	dfs(m, rc{coord.r, coord.c - 1}, visited)

	return true
}

type rc struct {
	r int
	c int
}
