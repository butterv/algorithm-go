package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var str string
	if _, err := fmt.Scanf("%s", &str); err != nil {
		panic(err)
	}

	digits := len(str)
	var result int64
	for i, s := range strings.Split(str, "") {
		n4, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}

		i := n4 * int64(math.Pow(float64(4), float64(digits-i-1)))
		result += i
	}

	fmt.Println(result)
}
