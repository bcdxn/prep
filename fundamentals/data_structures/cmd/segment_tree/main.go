package main

import (
	"fmt"

	ds "github.com/bcdxn/leetcode/fundamentals/data_structures"
)

func main() {
	nums := []int{7, 10, -3, 4, 2, 5}
	t := ds.NewSegmentTree(nums)

	sum := t.RangeSumQuery(0, 5)
	fmt.Println("sum: ", sum) // 25
	sum = t.RangeSumQuery(0, 3)
	fmt.Println("sum: ", sum) // 18
}
