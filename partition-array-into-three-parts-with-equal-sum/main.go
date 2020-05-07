package main

import (
	. "github.com/kilingzhang/leetcode"
	"time"
)

func main() {

	a := []int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}
	a = []int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}
	a = []int{13, -12, 23, -8, -10, 8, -2, 10, 14}
	a = []int{3, 17, 1, -9, 10, 0, -12, 20}
	a = []int{1, -1, 1, -1}

	Dump(canThreePartsEqualSum(a))

}

func canThreePartsEqualSum(a []int) bool {

	sum := 0
	for _, i := range a {
		sum += i
	}
	if sum%3 != 0 {
		return false
	} else {

		three := sum / 3
		i := 0
		j := len(a) - 1
		l := a[i]
		r := a[j]
		for ; ; {

			Dump(i, j, l, r, "---")

			if j <= i || i > (len(a)-1) || j <= 0 {
				break
			}

			if l == three && r == three {
				if j-i > 1 {
					return true
				} else {
					return false
				}
			}

			if l != three {
				i++
				l += a[i]
			}

			if r != three {
				j--
				r += a[j]
			}
			time.Sleep(time.Second)
		}

	}

	return false
}
