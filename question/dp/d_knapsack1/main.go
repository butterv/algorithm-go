package main

import (
	"fmt"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func count(W int, ws []int, vs []int) int {
	// Definition of DP
	dp := make([]int, W+1)

	for i := range ws {
		wi := ws[i]
		for j := W; j >= wi; j-- {
			dp[j] = max(dp[j], dp[j-wi]+vs[i])
		}
	}

	return dp[W]
}

// See: https://atcoder.jp/contests/dp/tasks/dp_d
func main() {
	var N, W int
	fmt.Scanf("%d %d", &N, &W)

	ws := make([]int, N)
	vs := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d", &ws[i], &vs[i])
	}

	fmt.Println(count(W, ws, vs))
}
