package main

import (
	"fmt"
	tool "github.com/kilingzhang/leetcode"
)

func productExceptSelf(nums []int) (ans []int) {
	l := len(nums)
	ldp := make([]int, l)
	rdp := make([]int, l)
	ldp[0] = nums[0]
	rdp[l-1] = nums[l-1]
	for i := 0; i < l; i++ {

		if i == 0 {
			ldp[i] = nums[i] * 1
		} else {
			ldp[i] = nums[i] * ldp[i-1]
		}

		if i == 0 {
			rdp[l-i-1] = nums[l-i-1] * 1
		} else {
			rdp[l-i-1] = nums[l-i-1] * rdp[l-i]
		}

	}

	fmt.Println(ldp, rdp)

	ans = make([]int, l)
	ans[0] = rdp[1]
	ans[l-1] = ldp[l-2]
	for i := 1; i < l-1; i++ {
		ans[i] = ldp[i-1] * rdp[i+1]
	}

	return ans
}

func main() {
	tool.Dump(productExceptSelf([]int{1, 2, 3, 4}))
}
