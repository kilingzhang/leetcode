package main

import tool "github.com/kilingzhang/leetcode"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	tool.Dump(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) (size int) {
	size = 0
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if size == 0 || dp[size-1] < nums[i] {
			dp[size] = nums[i]
			size++
		} else {
			x, y := 0, size-1
			for x < y {
				mid := (x + y) / 2
				if dp[mid] < nums[i] {
					x = mid + 1
				} else {
					y = mid
				}
			}
			dp[x] = nums[i]
		}
	}
	return size
}
