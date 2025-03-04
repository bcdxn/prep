package main

import (
	"fmt"

	ds "github.com/bcdxn/leetcode/fundamentals/data_structures"
)

func main() {
	edges := [][]int{
		{1, 2},
		{4, 1},
		{2, 4},
	}
	n := 4

	nodes := []int{}
	for i := range n {
		nodes = append(nodes, i+1)
	}

	uf := ds.NewUnionFind(nodes)

	for _, edge := range edges {
		n1, n2 := edge[0], edge[1]
		if !uf.Union(n1, n2) {
			fmt.Printf("%d <--> %d creates a cycle\n", n1, n2) // 2, 4
		}
	}
}
