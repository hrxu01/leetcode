package main

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = dp[0][i-1] + 1
	}

	for i := 1; i <= m; i++ {
		c1 := word1[i-1]
		for j := 1; j <= n; j++ {
			c2 := word2[j-1]
			if c1 == c2 {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[m][n]
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}