package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func ReadIntArray(sep string) ([]int, error) {
	var str string
	if sc.Scan() {
		str = sc.Text()
	}

	var nums []int
	for _, s := range strings.Split(str, sep) {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}

	return nums, nil
}
