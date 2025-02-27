package main

import "fmt"

// A medium problem that is quite simple IFF you can make the connection that the array of numbers
// represents a linked-list where each value in the array points to an index in the array (e.g. the
// _next_ node in the linked list). Then it's about finding the start of a cycle in the list using
// Floyd's algorithm.
func main() {
	nums := []int{1, 3, 4, 2, 2}
	dupe := findDuplicate(nums)
	fmt.Println("duplicate: ", dupe) // 2
}

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[0]

	for true {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow2 := nums[0]

	for slow != slow2 {
		slow = nums[slow]
		slow2 = nums[slow2]
	}

	return slow
}
