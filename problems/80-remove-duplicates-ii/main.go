package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := removeDuplicates(nums)
	fmt.Printf("k: %d -- %v", k, nums)
}

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	left := 2

	for i := 2; i < len(nums); i++ {
		if nums[left-2] != nums[i] {
			nums[left] = nums[i]
			left++
		}
	}

	return left
}
