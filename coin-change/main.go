package main

import (
	tool "github.com/kilingzhang/leetcode"
	"sort"
)

func main() {
	tool.Dump(
		coinChange(
			[]int{2},
			3,
		),
	)
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		var mixes []int
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			} else {
				if dp[i-coin] == -1 {
					continue
				}
				mixes = append(mixes, dp[i-coin])
			}
		}

		if len(mixes) == 0 {
			dp[i] = -1
			continue
		}

		mix := min(mixes)
		dp[i] = 1 + mix
	}
	return dp[amount]
}

func min(mixes []int) int {
	sort.Ints(mixes)
	return mixes[0]
}
