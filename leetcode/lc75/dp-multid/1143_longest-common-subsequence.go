package main

func longestCommonSubsequence(text1 string, text2 string) int {
	t1Len, t2Len := len(text1), len(text2)
	dp := make([][]int, t1Len+1)

	for i, _ := range dp {
		dp[i] = make([]int, t2Len+1)
	}

	for i := 1; i <= t1Len; i++ {
		for j := 1; j <= t2Len; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[t1Len][t2Len]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
