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
	if r < 0 || r >= rows || c < 0 || c >= cols {
		return 0
	}
	// blocked base case
	if m[r][c] == 1 {
		return 0
	}
	// already visited base case
	visitKey := fmt.Sprintf("%d:%d", r, c)
	if _, ok := visited[visitKey]; ok {
		return 0
	}
	// reached the end base case
	if r == rows-1 && c == cols-1 {
		return 1
	}

	// visiting new node
	visited[visitKey] = struct{}{}
	count := 0

	count += dfs(m, r+1, c, visited)
	count += dfs(m, r-1, c, visited)
	count += dfs(m, r, c+1, visited)
	count += dfs(m, r, c-1, visited)

	delete(visited, visitKey)

	return count
}
