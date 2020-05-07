package main

import (
	"fmt"
	tool "github.com/kilingzhang/leetcode"
)

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	maxSize := 0

	for i, bytes := range matrix {
		dp[i] = make([]int, len(bytes))
		for j, b := range bytes {
			if b-'0' == 1 {
				dp[i][j] = 1
				maxSize = 1
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {

			if dp[i][j] == 1 && dp[i][j-1] > 0 && dp[i-1][j] > 0 && dp[i-1][j-1] > 0 {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}

			if dp[i][j] > maxSize {
				maxSize = dp[i][j]
			}
		}
	}

	fmt.Println(dp)

	return maxSize * maxSize
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
	tool.Dump(maximalSquare([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
}
