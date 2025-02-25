package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{1, 2, 3}
	res := permute(nums)

	fmt.Println(res)
}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	subPerms := permute(nums[1:])
	res := [][]int{}

	for _, perm := range subPerms {
		for i := range len(perm) + 1 {
			newPerm := slices.Insert(perm, i, nums[0])
			res = append(res, newPerm)
		}
	}

	return res
}
