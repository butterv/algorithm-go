package main

import (
	"fmt"

	"github.com/istsh/algorithm-go/util"
)

func insertionSort(nums []int) {
	// start from index 1
	for i := 1; i < len(nums); i++ {
		numi := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > numi {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = numi
	}
}

// echo 5 1 10 4 7 3 6 2 8 9 | go run ./main.go
func main() {
	nums, err := util.ReadIntArray(" ")
	if err != nil {
		panic(err)
	}

	insertionSort(nums)
	fmt.Println(nums)
}
