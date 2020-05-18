package main

import (
	tool "github.com/kilingzhang/leetcode"
)

func maxProduct(nums []int) int {
	max := nums[0]
	l := len(nums)
	maxdp := max
	mindp := max
	for i := 1; i < l; i++ {

		maxdp, mindp = compare(nums[i], nums[i]*maxdp, nums[i]*mindp)

		if maxdp > max {
			max = maxdp
		}

	}
	return max
}

func compare(nums ...int) (int, int) {
	max := nums[0]
	min := nums[0]

	for i := 1; i < 3; i++ {
		if nums[i] > max {
			max = nums[i]
		}

		if nums[i] < min {
			min = nums[i]
		}
	}

	return max, min
}

func main() {
	tool.Dump(maxProduct([]int{0, 2}))
}
