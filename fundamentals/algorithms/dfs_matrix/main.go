package main

import "fmt"

func main() {
	m := [][]int{
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 1},
		{0, 1, 0, 0},
	}

	visited := make(map[rc]struct{})
	paths := dfs(m, 0, 0, visited)
	fmt.Println("num paths: ", paths)
}

type rc struct {
	r int
	c int
}

func dfs(m [][]int, r int, c int, visited map[rc]struct{}) int {
	rows := len(m)
	if rows < 1 {
		return 0
	}

	cols := len(m[0])
	if cols < 1 {
		return 0
	}

	// out of bounds
	if r < 0 || r >= rows || c < 0 || c >= cols {
		return 0
	}
	// already visited
	if _, ok := visited[rc{r, c}]; ok {
		return 0
	}
	// blocked
	if m[r][c] == 1 {
		return 0
	}
	// reached destination
	if r == rows-1 && c == cols-1 {
		return 1
	}

	count := 0
	visited[rc{r, c}] = struct{}{}

	count += dfs(m, r-1, c, visited)
	count += dfs(m, r+1, c, visited)
	count += dfs(m, r, c-1, visited)
	count += dfs(m, r, c+1, visited)

	delete(visited, rc{r, c})

	return count
}
