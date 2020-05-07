package main

import . "github.com/kilingzhang/leetcode"

func main() {

	var (
		candies   = 10
		numPeople = 3
	)

	peoples := distributeCandies(candies, numPeople)

	Dump(peoples)

}

func distributeCandies(candies int, numPeople int) []int {

	peoples := make([]int, numPeople)

	nums := getNums(candies, numPeople)

	for i := 1; i <= nums; i++ {

		for j := 0; j < numPeople; j++ {
			k := (j + 1) + numPeople*(i-1)
			if candies-k >= 0 {
				peoples[j] += k
				candies -= k
			} else {
				peoples[j] += candies
				candies = 0
			}
		}
	}

	return peoples
}

func getNums(candies int, numPeople int) int {

	sum := 0
	for i := 0; ; i++ {
		k := (numPeople+1)/2*numPeople + numPeople*numPeople*(i-1)
		if sum+k > candies {
			return i
		}

	}
}
