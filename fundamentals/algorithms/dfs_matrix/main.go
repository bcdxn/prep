package main

import "fmt"

func main() {
	matrix := [][]int{
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 1},
		{0, 1, 0, 0},
	}

	visited := map[rc]struct{}{}
	numPaths := dfs(matrix, 0, 0, visited)

	fmt.Println("numPaths: ", numPaths)
}

func dfs(m [][]int, r, c int, visited map[rc]struct{}) int {
	rows := len(m)
	cols := len(m[0])

	// at the destination
	if r == rows-1 && c == cols-1 {
		return 1
	}

	// already visited, out of bounds, or blocked
	if _, ok := visited[rc{r, c}]; ok || r < 0 || r >= rows || c < 0 || c >= cols || m[r][c] == 1 {
		return 0
	}

	// none of the above; mark the current node visited and keep recursing
	visited[rc{r, c}] = struct{}{}
	count := 0

	count += dfs(m, r-1, c, visited)
	count += dfs(m, r, c+1, visited)
	count += dfs(m, r+1, c, visited)
	count += dfs(m, r, c-1, visited)

	// unvisit the current node so that it can be used on other distinct paths
	delete(visited, rc{r, c})

	return count
}

type rc struct {
	r int
	c int
}
