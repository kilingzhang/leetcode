package main

import (
	. "github.com/kilingzhang/leetcode"
	"math"
)

func main() {

	target := 98160
	Dump(findContinuousSequence(target))

}

func findContinuousSequence(target int) (ret [][]int) {

	ret = make([][]int, 0)
	for i := int(math.Sqrt(float64(2 * target))); i >= 2; i-- {
		judge := target - i*(i-1)/2
		if judge%i == 0 {
			begin := judge / i
			temp := make([]int, 0)
			for j := 0; j < i; j++ {
				temp = append(temp, begin+j)
			}
			ret = append(ret, temp)
		}
	}

	return ret
}
