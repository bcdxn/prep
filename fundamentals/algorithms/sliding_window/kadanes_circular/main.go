package main

import "fmt"

func main() {
	nums1 := []int{4, -1, 2, -7, 3, 4}
	nums2 := []int{4, -1, 2, -7}

	sum := maxSum(nums1)
	fmt.Printf("sum %d\n", sum)
	sum = maxSum(nums2)
	fmt.Printf("sum %d\n", sum)
}

func maxSum(nums []int) int {
	totalSum := 0
	maxSum := nums[0]
	minSum := nums[0]
	currMaxSum := 0
	currMinSum := 0
	currR := 0

	for currR < len(nums) {
		if currMaxSum < 0 {
			currMaxSum = 0
		}

		if currMinSum > 0 {
			currMinSum = 0
		}

		currMinSum += nums[currR]
		currMaxSum += nums[currR]
		totalSum += nums[currR]

		if currMaxSum > maxSum {
			maxSum = currMaxSum
		}

		if currMinSum < minSum {
			minSum = currMinSum
		}

		currR += 1
	}

	if maxSum > 0 {
		notMinSum := totalSum - minSum
		if maxSum > notMinSum {
			return maxSum
		} else {
			return notMinSum
		}
	} else {
		return maxSum
	}
}
