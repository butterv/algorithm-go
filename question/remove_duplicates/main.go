package main

import (
	"fmt"
	"sort"
)

// Method to remove duplicates from array.
// nums from 1 to 100.

func removeDuplicates(nums []int) []int {
	sort.Ints(nums)

	previous := nums[0]
	result := make([]int, len(nums))
	result[0] = previous

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if previous != num {
			result[i] = num
		}
		previous = num
	}

	return result
}

func main() {
	data := [][]int{
		{1, 1, 2, 2, 3, 4, 5},
		{1, 1, 1, 1, 1, 1, 1},
		{1, 2, 3, 4, 5, 6, 7},
		{1, 2, 1, 1, 1, 1, 1},
	}

	for _, d := range data {
		result := removeDuplicates(d)
		fmt.Println(result)
	}
}
