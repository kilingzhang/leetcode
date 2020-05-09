package main

import tool "github.com/kilingzhang/leetcode"

func pivotIndex(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	left := 0
	for i := 0; i < len(nums); i++ {
		if i != 0{
			left += nums[i-1]
		}
		right := sum - left - nums[i]
		if left == right {
			return i
		}
	}

	return -1
}

func main() {
	tool.Dump(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}
