package main

import tool "github.com/kilingzhang/leetcode"

func waysToStep(n int) int {

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	if n == 3 {
		return 4
	}

	dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 2
	dp[2] = 4

	for i := 3; i < n; i++ {
		dp[i] = (dp[i-1]%1000000007 + dp[i-2]%1000000007 + dp[i-3]%1000000007) % 1000000007
	}

	return dp[n-1]
}

func main() {
	tool.Dump(waysToStep(76))
}
