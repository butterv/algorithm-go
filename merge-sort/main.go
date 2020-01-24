package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

// echo 9 1 5 6 2 8 10 3 7 4 | go run ./main.go
func main() {
	nums, err := read()
	if err != nil {
		panic(err)
	}

	fmt.Println(mergeSort(nums))
}

func read() ([]int, error) {
	var str string
	if sc.Scan() {
		str = sc.Text()
	}

	var nums []int
	for _, s := range strings.Split(str, " ") {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}

	return nums, nil
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	middle := len(nums) / 2
	//	# ここで分割を行う
	left := nums[:middle]
	right := nums[middle:]

	//	# 再帰的に分割を行う
	left = mergeSort(left)
	right = mergeSort(right)

	// # returnが返ってきたら、結合を行い、結合したものを次に渡す
	return merge(left, right)
}

func merge(left, right []int) []int {
	var merged []int

	li, ri := 0, 0

	//	# ソート済み配列をマージするため、それぞれ左から見ていくだけで良い
	for li < len(left) && ri < len(right) {
		// # ここで=をつけることで安定性を保っている
		if left[li] <= right[ri] {
			merged = append(merged, left[li])
			li++
		} else {
			merged = append(merged, right[ri])
			ri++
		}
	}

	//	# 上のwhile文のどちらかがfalseになった場合終了するため、あまりをextendする
	if li < len(left) {
		merged = append(merged, left[li:]...)
	}
	if ri < len(right) {
		merged = append(merged, right[ri:]...)
	}

	return merged
}
