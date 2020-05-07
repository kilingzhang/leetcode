package main

import (
	"fmt"
	tool "github.com/kilingzhang/leetcode"
)

func rob(nums []int) int {

	numCount := len(nums)

	if numCount <= 0 {
		return 0
	}

	if numCount == 1 {
		return nums[0]
	}

	if numCount == 2 {
		return max(nums[0], nums[1])
	}

	if numCount == 3 {
		return max(nums[0]+nums[2], nums[1])
	}

	dp := make([]int, numCount+1)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	dp[2] = max(nums[0]+nums[2], nums[1])

	for i := 3; i < len(nums); i++ {
		d := max(dp[i-2]+nums[i], dp[i-3]+nums[i])
		d = max(d, dp[i-1])
		dp[i] = d
	}

	fmt.Println(dp)
	return dp[numCount-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	tool.Dump(rob([]int{2, 1, 1, 3}))
}
