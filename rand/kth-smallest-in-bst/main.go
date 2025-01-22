package main

import (
	"fmt"

	ds "github.com/bcdxn/leetcode/fundamentals/data_structures"
)

func main() {
	input := []int{4, 3, 5, 2}

	tree := ds.NewBST[int]()

	for _, val := range input {
		tree.Insert(val)
	}

	fmt.Println("tree: ", tree.StringInOrder())
	fmt.Println("4th smallest element", findKthSmallest(tree, 4))
}

var count = 0

func findKthSmallest(tree *ds.BST[int], k int) int {
	count := k
	res := 0

	var dfs func(root *ds.BSTNode[int])
	dfs = func(root *ds.BSTNode[int]) {
		if root == nil {
			return
		}

		dfs(root.Left)
		count -= 1
		if count == 0 {
			res = root.Item
			return
		}
		dfs(root.Right)
	}

	dfs(tree.Root)
	return res
}
