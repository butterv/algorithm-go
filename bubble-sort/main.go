package main

import (
	"fmt"

	"github.com/istsh/algorithm-go/util"
)

func bubbleSort(nums []int, length int) {
	for i := 0; i < length-1; i++ {
		for j := length - 1; j >= 1; j-- {
			if nums[j] < nums[j-1] {
				tmp := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = tmp
			}
		}
	}
}

// echo 9 1 5 6 2 8 10 3 7 4 | go run ./main.go
func main() {
	nums, err := util.ReadIntArray(" ")
	if err != nil {
		panic(err)
	}

	bubbleSort(nums, len(nums))
	fmt.Println(nums)
}
