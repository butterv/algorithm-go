package main

import "fmt"

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

func calc(N int, as []int) int {
	dp := make([]int, N)
	for i := 0; i < N; i++ {
		dp[i] = INF
	}

	dp[0] = 0

	for i := 1; i < N; i++ {
		chmin(&dp[i], dp[i-1]+abs(as[i]-as[i-1]))
		if i > 1 {
			chmin(&dp[i], dp[i-2]+abs(as[i]-as[i-2]))
		}
	}

	return dp[N-1]
}

// See: https://atcoder.jp/contests/abc040/tasks/abc040_c
func main() {
	var N int
	fmt.Scan(&N)
	as := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&as[i])
	}

	fmt.Println(calc(N, as))
}
