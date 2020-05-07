package main

import (
	tool "github.com/kilingzhang/leetcode"
	"math"
)

func main() {
	var prices []int = []int{7, 1, 5, 3, 6, 4}
	tool.Dump(maxProfit(prices))
}

func maxProfit(prices []int) int {

	max := 0

	t := math.MaxInt64
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] > t {
			continue
		}
		t = prices[i]
		tt := math.MinInt8
		for j := i + 1; j < len(prices); j++ {

			if prices[j] < tt {
				continue
			}

			tt = prices[j]

			if max < prices[j]-prices[i] {
				max = prices[j] - prices[i]
			}
		}
	}

	return max
}
