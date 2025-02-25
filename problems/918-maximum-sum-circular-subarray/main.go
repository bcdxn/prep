package main

import "fmt"

func main() {
	nums := []int{1, -2, 3, -2}
	maxSum := maxSubarraySumCircular(nums)
	fmt.Println("max sum: ", maxSum) // 3
	nums = []int{5, -3, 5}
	maxSum = maxSubarraySumCircular(nums)
	fmt.Println("max sum: ", maxSum) // 10
	nums = []int{-3, -2, -3}
	maxSum = maxSubarraySumCircular(nums)
	fmt.Println("max sum: ", maxSum) // -2
}

func maxSubarraySumCircular(nums []int) int {
	total := 0
	currMax := 0
	currMin := 0
	max := nums[0]
	min := nums[0]

	r := 0
	for r < len(nums) {
		if currMax < 0 {
			currMax = 0
		}

		if currMin > 0 {
			currMin = 0
		}

		total += nums[r]
		currMax += nums[r]
		currMin += nums[r]

		if currMax > max {
			max = currMax
		}

		if currMax < min {
			min = currMin
		}

		r++
	}

	if max < 0 {
		return max
	}

	if total-min > max {
		return total - min
	} else {
		return max
	}
}
