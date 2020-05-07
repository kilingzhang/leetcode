package main

import (
	tool "github.com/kilingzhang/leetcode"
)

func maxSubArray(nums []int) int {

	sum := 0
	max := nums[0]
	for _, num := range nums {

		if sum > 0 {
			sum += num
		} else {
			sum = num
		}

		if max < sum {
			max = sum
		}

	}

	return max
}

func main() {
	tool.Dump(maxSubArray([]int{3, 2, -3, -1, 1, -3, 1, -1}))
}
