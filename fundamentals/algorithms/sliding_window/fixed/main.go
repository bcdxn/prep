package main

import "fmt"

func main() {
	var nums = []int{8, 3, -2, 4, 5, -1, 0, 5, 3, 9, -6}
	var k = 5

	max := maxSum(nums, k)

	fmt.Println("max sum: ", max)
}

func maxSum(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}

	// calculate initial window sum
	max := 0
	for i := 0; i < k; i++ {
		max += nums[i]
	}
	// slide the window
	windowSum := max
	for i := 0; i < len(nums)-k; i++ {
		windowSum = windowSum - nums[i] + nums[i+k]
		if windowSum > max {
			max = windowSum
		}
	}

	return max
}
