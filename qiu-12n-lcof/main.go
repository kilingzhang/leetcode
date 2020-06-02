package main

import tool "github.com/kilingzhang/leetcode"

func sumNums(n int) int {
	ans := 0
	var sum func(int) bool
	sum = func(n int) bool {
		ans += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}

func main() {
	tool.Dump(sumNums(3))
}