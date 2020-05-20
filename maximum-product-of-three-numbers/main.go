package main

import (
	"fmt"
	"sort"
)

func maximumProduct(nums []int) int {
	l := len(nums)
	sort.Ints(nums)
	a := nums[0] * nums[1] * nums[l-1]
	b := nums[l-1] * nums[l-2] * nums[l-3]
	if a > b {
		return a
	}
	return b
}


func main() {
	fmt.Println(maximumProduct([]int{1, 2, 3, 4}))
}
