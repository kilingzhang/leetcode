package main

import tool "github.com/kilingzhang/leetcode"

func subarraySum(nums []int, k int) (ans int) {
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end >= 0; end-- {
			sum += nums[end]
			if sum == k {
				ans++
			}
		}
	}
	return ans
}

func main() {
	tool.Dump(subarraySum([]int{1, 1, 1}, 2))
}
