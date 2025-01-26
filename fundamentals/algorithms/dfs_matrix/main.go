package main

import "fmt"

func main() {
	matrix := [][]int{
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 1},
		{0, 1, 0, 0},
	}

	paths := dfs(matrix, 0, 0, make(map[rc]struct{}))
	fmt.Println("paths: ", paths)
}

func dfs(m [][]int, r int, c int, visited map[rc]struct{}) int {
	rows := len(m)
	cols := len(m[0])

	if r == rows-1 && c == cols-1 {
		return 1
	}

	count := 0

	if _, ok := visited[rc{r - 1, c}]; !ok && r-1 >= 0 && m[r-1][c] != 1 {
		visited[rc{r - 1, c}] = struct{}{}
		count += dfs(m, r-1, c, visited)
		delete(visited, rc{r - 1, c})
	}

	if _, ok := visited[rc{r + 1, c}]; !ok && r+1 < rows && m[r+1][c] != 1 {
		visited[rc{r + 1, c}] = struct{}{}
		count += dfs(m, r+1, c, visited)
		delete(visited, rc{r + 1, c})
	}

	if _, ok := visited[rc{r, c - 1}]; !ok && c-1 >= 0 && m[r][c-1] != 1 {
		visited[rc{r, c - 1}] = struct{}{}
		count += dfs(m, r, c-1, visited)
		delete(visited, rc{r, c - 1})
	}

	if _, ok := visited[rc{r, c + 1}]; !ok && c+1 < cols && m[r][c+1] != 1 {
		visited[rc{r, c + 1}] = struct{}{}
		count += dfs(m, r, c+1, visited)
		delete(visited, rc{r, c + 1})
	}

	return count
}

type rc struct {
	r int
	c int
}
