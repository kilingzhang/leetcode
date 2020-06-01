package main

import (
	"fmt"
	"math"
)

func kidsWithCandies(candies []int, extraCandies int) (ans []bool) {
	max := math.MinInt64
	for i := 0; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}
	ans = make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		ans[i] = extraCandies+candies[i] >= max
	}
	return ans
}

func main() {
	fmt.Println(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1))
}
