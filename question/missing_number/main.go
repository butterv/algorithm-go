package main

import (
	"fmt"

	"github.com/istsh/algorithm-go/util"
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
	nums, err := util.ReadIntArray(" ")
	if err != nil {
		panic(err)
	}

	var count int
	fmt.Scanf("%d", &count)

	if len(nums) > count {
		panic("invalid args")
	}

	missings := getMissingNumber(nums, count)
	fmt.Println(missings)
}
