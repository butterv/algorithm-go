package main

import (
	"fmt"
)

// Finds missing numbers in a given array of integers.
// nums from 1 to 100.

func getMissingNumber(nums []int, count int) []int {
	m := map[int]struct{}{}
	for _, num := range nums {
		m[num] = struct{}{}
	}

	var missings []int
	for i := 1; i <= count; i++ {
		if _, ok := m[i]; !ok {
			missings = append(missings, i)
		}
	}

	return missings
}

func main() {
	data := []struct {
		nums  []int
		count int
	}{
		{
			nums:  []int{1, 2, 3, 4, 6},
			count: 6,
		},
		{
			nums:  []int{1, 2, 3, 4, 6, 7, 9, 8, 10},
			count: 10,
		},
		{
			nums:  []int{1, 2, 3, 4, 6, 9, 8},
			count: 10,
		},
	}

	for _, d := range data {
		if len(d.nums) > d.count {
			panic("invalid args")
		}

		missings := getMissingNumber(d.nums, d.count)
		fmt.Println(missings)
	}
}
