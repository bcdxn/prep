package main

import "fmt"

func main() {
	nums := []int{8, 3, -2, 4, 5, -1, 0, 5, 3, 9, -6}
	res := maxContiguousSegment(nums, 5)
	fmt.Println("max:", res) // 18
}

func maxContiguousSegment(nums []int, k int) int {
	max := 0

	if k >= len(nums) {
		for _, val := range nums {
			max += val
		}
		return max
	}

	// get sum in initial window
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}
	max = windowSum

	// slide the window
	for i := 0; i < len(nums)-k; i++ {
		windowSum = windowSum - nums[i] + nums[i+k]
		if windowSum > max {
			max = windowSum
		}
	}

	return max
}
