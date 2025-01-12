package main

import "fmt"

func main() {
	nums := []int{8, 3, -2, 4, 5, -1, 0, 5, 3, 9, -6}
	res := maxSumInWindow(nums, 5)
	fmt.Println("max:", res) // 18
}

func maxSumInWindow(nums []int, k int) int {
	if k > len(nums) {
		return 0
	}
	windowSum := 0

	// sum initial window
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}
	max := windowSum

	// slide window and update sum
	for i := 0; i < len(nums)-k; i++ {
		windowSum = windowSum - nums[i] + nums[i+k]
		if windowSum > max {
			max = windowSum
		}
	}

	return max
}
