package main

import (
	"fmt"
)

func chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}

func count(N int, abc [][3]int) int {
	// Definition of DP
	dp := make([][3]int, N+1)

	for i := 0; i < N; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if j == k {
					continue
				}
				chmax(&dp[i+1][k], dp[i][j]+abc[i][k])
			}
		}
	}

	var result int
	for i := 0; i < 3; i++ {
		chmax(&result, dp[N][i])
	}

	return result
}

// See: https://atcoder.jp/contests/dp/tasks/dp_c
func main() {
	var N int
	fmt.Scan(&N)

	abc := make([][3]int, N)
	for i := 0; i < N; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&abc[i][j])
		}
	}

	fmt.Println(count(N, abc))
}
