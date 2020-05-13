package main

import tool "github.com/kilingzhang/leetcode"

func missingNumber(nums []int) (ans int) {
	l := len(nums)
	ans = l
	for i := 0; i < l; i++ {
		ans ^= nums[i] ^ i
	}
	return ans
}

func main() {
	tool.Dump(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
