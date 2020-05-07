package main

import tool "github.com/kilingzhang/leetcode"

func main() {
	nums := []int{2, 4, -2, -3}
	tool.Dump(increasingTriplet(nums))
}

func increasingTriplet(nums []int) bool {

	size := 0
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

	if size >= 3 {
		return true
	}

	return false
}
