package main

import (
	tool "github.com/kilingzhang/leetcode"
)

func jump(nums []int) int {

	step := 0
	currentPos := 0
	distance := len(nums) - 1

	if distance == 0 {
		return step
	}

	for distance != 0 {

		step++
		skip := nums[currentPos] + currentPos + 1

		if skip > distance {
			return step
		}

		canJumps := nums[currentPos+1 : skip]

		max := canJumps[0]
		pos := 0

		for i, canJump := range canJumps {
			if max+pos <= canJump+i {
				max = canJump
				pos = i
			}
		}

		currentPos += pos + 1
	}

	return step
}

func main() {
	tool.Dump(jump([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0}))
}
