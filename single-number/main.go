package main

import tool "github.com/kilingzhang/leetcode"

func singleNumber(nums []int) (ans int) {
	ans = nums[0]
	if len(nums) > 1 {
		for i := 1; i < len(nums); i++ {
			ans = ans ^ nums[i]
		}
	}
	return ans
}

func main() {
	tool.Dump(singleNumber([]int{4, 1, 1, 2, 2, 3, 3}))
}
