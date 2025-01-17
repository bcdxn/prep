package main

import (
	"fmt"

	ds "github.com/bcdxn/leetcode/fundamentals/data_structures"
)

func main() {
	// _ := 5
	preReqs := [][]int{
		{0, 1},
		{0, 2},
		{1, 3},
		// {1, 4}, // uncomment to allow completion
		{3, 4},
		{4, 1}, // comment out to allow completion
	}

	matrix := ds.ParseMatrixFromEdges[int](preReqs)

	cycle := containsCycle(matrix)

	if cycle {
		fmt.Println("cannot complete courses")
	} else {
		fmt.Println("can complete courses")
	}
}

func containsCycle(matrix map[int][]int) bool {
	for node := range matrix {
		cycle := dfs(node, matrix, make(map[int]struct{}))
		if cycle {
			return true
		}
	}

	return false
}

func dfs(node int, matrix map[int][]int, visited map[int]struct{}) bool {
	if _, ok := visited[node]; ok {
		return true
	}
	visited[node] = struct{}{}

	cycle := false

	for i := range matrix[node] {
		cycle = cycle || dfs(matrix[node][i], matrix, visited)
	}

	delete(visited, node)

	return cycle
}
