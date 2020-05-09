package main

import (
	tool "github.com/kilingzhang/leetcode"
)

func dominantIndex(nums []int) int {

	first := 0
	second := 0
	index := -1
	for i, num := range nums {

		if first < num {
			index = i
			second = first
			first = num
		}

		if second < num && num < first {
			second = num
		}

	}

	if second == 0 || first/second >= 2 {
		return index
	}

	return -1
}

func main() {
	tool.Dump(dominantIndex([]int{0, 0, 3, 2}))
}
