package main

import (
	tool "github.com/kilingzhang/leetcode"
)

func myPow(x float64, n int) float64 {

	if x == 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	if n == 1 {
		return x
	}

	if n%2 == 0 {
		ans := myPow(x, n/2)
		return ans * ans
	} else {
		n--
		ans := myPow(x, n/2)
		return ans * ans * x
	}

}

func main() {
	tool.Dump(myPow(0.00001, 2147483647))
}
