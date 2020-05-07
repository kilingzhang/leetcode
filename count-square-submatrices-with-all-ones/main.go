package main

import tool "github.com/kilingzhang/leetcode"

func countSquares(matrix [][]int) int {

	dp := make([][]int, len(matrix))
	for i, ants := range matrix {
		dp[i] = make([]int, len(ants))
		for j, n := range ants {
			if n == 1 {
				dp[i][j] = 1
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {

			if dp[i][j] == 1 && dp[i][j-1] > 0 && dp[i-1][j] > 0 && dp[i-1][j-1] > 0 {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
			}

		}
	}

	sum := 0
	for _, ints := range dp {
		for _, count := range ints {
			sum += count
		}
	}

	return sum
}

func min(a, b, c int) int {
	return _min(_min(a, b), c)
}
func _min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	tool.Dump(countSquares([][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}}))
}
