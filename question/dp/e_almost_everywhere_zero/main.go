package main

import (
	"fmt"
)

func counter(s string, K int) int {
	// DP定義
	var dp [105][4][2]int
	var n int = len(s)
	dp[0][0][0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 2; k++ {
				nd := int(s[i] - '0')
				for d := 0; d < 10; d++ {
					ni, nj, nk := i+1, j, k
					if d != 0 {
						nj++
					}
					if nj > K {
						continue
					}
					if k == 0 {
						if d > nd {
							continue
						}
						if d < nd {
							nk = 1
						}
					}
					dp[ni][nj][nk] += dp[i][j][k]
				}
			}
		}
	}
	var ans int = dp[n][K][0] + dp[n][K][1]
	return ans
}

// See: https://atcoder.jp/contests/abc154/tasks/abc154_e
func main() {
	var N string
	var K int
	fmt.Scan(&N)
	fmt.Scan(&K)
	fmt.Println(counter(N, K))
}
