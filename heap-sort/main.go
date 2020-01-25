package main

import (
	"fmt"

	"github.com/istsh/algorithm-go/util"
)

func heapSort(nums []int) {
	length := len(nums)

	for i := length/2 - 1; i >= 0; i-- {
		heap(nums, length, i)
	}

	for i := length - 1; i >= 0; i-- {
		if nums[0] > nums[i] {
			tmp := nums[0]
			nums[0] = nums[i]
			nums[i] = tmp

			heap(nums, i-1, 0)
		}
	}
}

func heap(nums []int, length, root int) {
	largest := root
	left := root*2 + 1
	right := root*2 + 2

	if left < length && nums[left] > nums[largest] {
		largest = left
	}

	if right < length && nums[right] > nums[largest] {
		largest = right
	}

	if largest != root {
		swap := nums[root]
		nums[root] = nums[largest]
		nums[largest] = swap

		heap(nums, length, largest)
	}
}

// echo 9 1 5 6 2 8 10 3 7 4 | go run ./main.go
func main() {
	nums, err := util.ReadIntArray(" ")
	if err != nil {
		panic(err)
	}

	heapSort(nums)
	fmt.Println(nums)
}
