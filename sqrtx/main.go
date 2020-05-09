package main

import (
	tool "github.com/kilingzhang/leetcode"
)

func mySqrt(x int) int {

	left := 0
	right := x

	ans := 0
	for left <= right {

		mid := left + (right-left)/2

		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
			ans = mid
		}

	}
	return ans
}

func main() {
	tool.Dump(mySqrt(8))
}
