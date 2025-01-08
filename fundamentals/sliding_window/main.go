package main

import "fmt"

func main() {
	nums := []int{8, 3, -2, 4, 5, -1, 0, 5, 3, 9, -6}
	res := maxContiguousSegment(nums, 5)
	fmt.Println("max:", res)
}

func maxContiguousSegment(nums []int, k int) int {
	max := 0 // track maximum window sum

	if k > len(nums) {
		for i := 0; i < len(nums); i++ {
			max += nums[i]
		}
		return max
	}

	for i := 0; i < k; i++ {
		max += nums[i]
	}

	sum := max // track current window

	for i := 0; i < len(nums)-k; i++ {
		sum = sum - nums[i] + nums[i+k]
		if sum > max {
			max = sum
		}
	}

	return max
}
