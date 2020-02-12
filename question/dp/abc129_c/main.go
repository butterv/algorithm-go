package main

import "fmt"

func calc(N int, am map[int]struct{}) int64 {
	dp := make([]int64, N+1)
	dp[0] = 1
	if _, ok := am[1]; !ok {
		dp[1] = 1
	}

	for i := 2; i <= N; i++ {
		oneBefore, twoBefore := dp[i-1], dp[i-2]
		if oneBefore == 0 && twoBefore == 0 {
			break
		}

		if _, ok := am[i]; ok {
			dp[i] = 0
		} else {
			dp[i] = (oneBefore + twoBefore) % 1000000007
		}
	}

	return dp[N]
}

// See: https://atcoder.jp/contests/abc129/tasks/abc129_c
func main() {
	var N, M int
	fmt.Scanf("%d %d", &N, &M)
	am := map[int]struct{}{}
	for i := 1; i <= M; i++ {
		var a int
		fmt.Scan(&a)
		am[a] = struct{}{}
	}
	fmt.Println(calc(N, am) % 1000000007)
}
