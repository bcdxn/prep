package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	top := kMostFrequent(nums, 2)
	fmt.Println("top 2 elements: ", top)
}

func kMostFrequent(nums []int, k int) []int {
	set := map[int]int{}
	// get count of occurences of each number in the list
	for _, val := range nums {
		if count, ok := set[val]; ok {
			set[val] = count + 1
		} else {
			set[val] = 1
		}
	}
	// create a bucket to 'bucket sort' by count
	buckets := make([][]int, len(nums)+1)
	// add each number in the set to the index that corresponds to the count of occurences
	for num, count := range set {
		if buckets[count] == nil {
			buckets[count] = []int{}
		}
		buckets[count] = append(buckets[count], num)
	}
	top := make([]int, 0, k)
	// return k most frequent
	for i := len(buckets) - 1; i > 0 && k > 0; i-- {
		for j := 0; j < len(buckets[i]) && k > 0; j++ {
			top = append(top, buckets[i][j])
			k -= 1
		}
	}
	return top
}
