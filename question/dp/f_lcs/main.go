package main

import (
	"fmt"
)

func makeDp(iLen, jLen int) [][]int {
	dp := make([][]int, iLen)
	for i := range dp {
		dp[i] = make([]int, jLen)
	}
	return dp
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func reverse(bs []byte) {
	var beforeHarfStart, afterHarfStart, loopCnt int
	length := len(bs)
	if length%2 == 0 {
		beforeHarfStart = length/2 - 1
		afterHarfStart = length / 2
		loopCnt = length / 2
	} else {
		beforeHarfStart = (length - 1) / 2
		afterHarfStart = (length - 1) / 2
		loopCnt = (length + 1) / 2
	}

	for i := 0; i < loopCnt; i++ {
		s1 := bs[beforeHarfStart-i]
		s2 := bs[afterHarfStart+i]
		bs[beforeHarfStart-i] = s2
		bs[afterHarfStart+i] = s1
	}
}

func extract(s, t string) string {
	sLen, tLen := len(s), len(t)
	// Definition of DP
	//var dp [3001][3001]int
	dp := makeDp(sLen+1, tLen+1)

	for i := 0; i < sLen; i++ {
		for j := 0; j < tLen; j++ {
			if s[i] == t[j] {
				dp[i+1][j+1] = max(dp[i+1][j+1], dp[i][j]+1)
			}
			dp[i+1][j+1] = max(dp[i+1][j+1], dp[i+1][j])
			dp[i+1][j+1] = max(dp[i+1][j+1], dp[i][j+1])
		}
	}

	var result []byte
	i, j := sLen, tLen
	for i > 0 && j > 0 {
		if dp[i][j] == dp[i-1][j] {
			// (i-1, j) -> (i, j) と更新されていた場合
			i-- // DP の遷移を遡る
		} else if dp[i][j] == dp[i][j-1] {
			// (i, j-1) -> (i, j) と更新されていた場合
			j-- // DP の遷移を遡る
		} else {
			// (i-1, j-1) -> (i, j) と更新されていた場合
			result = append(result, s[i-1]) // このとき s[i-1] == t[j-1] なので、t[j-1] + res でも OK
			i--                             // DP の遷移を遡る
			j--
		}
	}

	reverse(result)
	return string(result)
}

// Ses: https://atcoder.jp/contests/dp/tasks/dp_f
func main() {
	var s, t string
	fmt.Scan(&s)
	fmt.Scan(&t)

	fmt.Println(extract(s, t))
}
