package main

import (
	tool "github.com/kilingzhang/leetcode"
)

func plusOne(digits []int) []int {

	i := 0
	j := len(digits) - 1
	for i < j {
		digits[i], digits[j] = digits[j], digits[i]
		i++
		j--
	}

	isPlus := false
	for i, _ := range digits {
		sum := digits[i] + 1
		if sum < 10 {
			digits[i] = sum % 10
			isPlus = false
			break
		}
		digits[i] = sum % 10
		isPlus = true
	}

	if isPlus {
		digits = append(digits, 1)
	}

	i = 0
	j = len(digits) - 1
	for i < j {
		digits[i], digits[j] = digits[j], digits[i]
		i++
		j--
	}

	return digits
}

func main() {
	tool.Dump(plusOne([]int{9}))
}
