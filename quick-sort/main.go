package main

import (
	"fmt"

	"github.com/istsh/algorithm-go/util"
)

func partition(nums []int, p, r int) int {
	x := nums[r]
	i := p - 1
	for j := p; j <= r-1; j++ {
		if nums[j] <= x {
			i = i + 1
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
		}
	}
	tmp := nums[i+1]
	nums[i+1] = nums[r]
	nums[r] = tmp
	return i + 1
}

func quickSort(nums []int, p, r int) {
	if p < r {
		q := partition(nums, p, r)
		quickSort(nums, p, q-1)
		quickSort(nums, q+1, r)
	}
}

// echo 9 1 5 6 2 8 10 3 7 4 | go run ./main.go
func main() {
	nums, err := util.ReadIntArray(" ")
	if err != nil {
		panic(err)
	}

	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
