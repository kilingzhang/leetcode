package main

import (
	tool "github.com/kilingzhang/leetcode"
)

func climbStairs(n int) int {

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	return climbStairs(n-1) + climbStairs(n-2)
}

func main() {
	tool.Dump(climbStairs(44))
}
