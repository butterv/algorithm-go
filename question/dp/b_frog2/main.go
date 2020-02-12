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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func calcCost(N, K int, hs []int) int {
	// Definition of DP
	dp := make([]int, N+K)
	// Initialization
	for i := 0; i < N+K; i++ {
		dp[i] = INF
	}
	// Index zero
	dp[0] = 0

	for i := 0; i < N; i++ {
		for j := 1; j <= K; j++ {
			chmin(&dp[i+j], dp[i]+abs(hs[i]-hs[i+j]))
		}
	}

	return dp[N-1]
}

// See: https://atcoder.jp/contests/dp/tasks/dp_b
func main() {
	var N, K int
	fmt.Scanf("%d %d", &N, &K)

	hs := make([]int, N+K)
	for i := 0; i < N; i++ {
		fmt.Scan(&hs[i])
	}

	fmt.Println(calcCost(N, K, hs))
}
