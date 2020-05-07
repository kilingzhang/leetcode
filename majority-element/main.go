package main

import (
	tool "github.com/kilingzhang/leetcode"
	"math"
)

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	tool.Dump(majorityElement(nums))
}

func majorityElement(nums []int) int {
	counts := map[int]int{}
	for _, num := range nums {
		counts[num] ++
	}
	if i, count := max(counts); count > len(nums)/2 {
		return i
	}

	return -1
}

func max(nums map[int]int) (index int, max int) {
	max = math.MinInt32

	for i, num := range nums {
		if num > max {
			index = i
			max = num
		}
	}

	return index, max
}
