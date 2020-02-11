package main

import (
	"fmt"
)

const INF = 1001001001

func chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

func chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func calcCost(N int, hs []int) int {
	// Definition of DP
	dp := make([]int, N)
	// Initialization
	for i := 0; i < N; i++ {
		dp[i] = INF
	}
	// Index zero
	dp[0] = 0

	for i := 1; i < N; i++ {
		chmin(&dp[i], dp[i-1]+abs(hs[i]-hs[i-1]))
		if i > 1 {
			chmin(&dp[i], dp[i-2]+abs(hs[i]-hs[i-2]))
		}
	}

	return dp[N-1]
}

// See: https://atcoder.jp/contests/dp/tasks/dp_a
func main() {
	var N int
	fmt.Scan(&N)

	hs := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&hs[i])
	}

	fmt.Println(calcCost(N, hs))
}
