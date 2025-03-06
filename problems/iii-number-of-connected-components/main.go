package main

import "fmt"

func main() {
	n := 6
	edges := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{4, 5},
	}

	count := countComponents(n, edges)
	fmt.Println("count of components:", count) // 2
}

func countComponents(n int, edges [][]int) int {
	count := n
	uf := newUnionFind(n)

	for _, edge := range edges {
		if uf.Union(edge[0], edge[1]) {
			count -= 1
		}
	}

	return count
}

// create a new UnionFind data structure
func newUnionFind(n int) UnionFind {
	parents := make([]int, n)
	ranks := make([]int, n)
	for i := range n {
		parents[i] = i
	}

	return UnionFind{
		parents,
		ranks,
	}
}

// Find the root parent of a given node (and perform path compression along the way)
func (uf *UnionFind) Find(n int) int {
	node := n

	for uf.parents[node] != node {
		uf.parents[node] = uf.parents[uf.parents[node]]
		node = uf.parents[node]
	}

	return node
}

// Union and return the root parent
func (uf *UnionFind) Union(n1, n2 int) bool {
	p1, p2 := uf.Find(n1), uf.Find(n2)

	if p1 == p2 {
		return false
	}

	if uf.ranks[p1] > uf.ranks[p2] {
		uf.parents[p2] = p1
	} else if uf.ranks[p2] > uf.ranks[p1] {
		uf.parents[p1] = p2
	} else {
		uf.parents[p2] = p1
		uf.ranks[p1] += 1
	}

	return true
}

// UnionFind struct
type UnionFind struct {
	parents []int
	ranks   []int
}
