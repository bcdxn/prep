package main

import "fmt"

func main() {
	m := [][]int{
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 1},
		{0, 1, 0, 0},
	}

	paths := dfs(m, 0, 0, make(map[string]struct{}))
	fmt.Println("paths:", paths) // 2 paths
}

func dfs(m [][]int, r int, c int, visited map[string]struct{}) int {
	rows := len(m)
	if rows < 1 {
		return 0
	}
	cols := len(m[0])
	if cols < 1 {
		return 0
	}

	// out of bounds base case
	if r < 0 || c < 0 || r >= rows || c >= cols {
		return 0
	}
	// already visited base case
	visitKey := fmt.Sprintf("%d:%d", r, c)
	if _, ok := visited[visitKey]; ok {
		return 0
	}
	// path complete base case
	if r == rows-1 && c == cols-1 {
		return 1
	}

	// mark node as visited
	visited[visitKey] = struct{}{}

	count := 0
	count += dfs(m, r+1, c, visited)
	count += dfs(m, r-1, c, visited)
	count += dfs(m, r, c+1, visited)
	count += dfs(m, r, c-1, visited)

	return count
}
