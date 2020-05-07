package main

type NumArray struct {
	dp []int
}

func Constructor(nums []int) NumArray {
	n := new(NumArray)
	numCount := len(nums)
	n.dp = make([]int, numCount+1)

	if numCount <= 0 {
		n.dp[0] = 0
		return *n
	}

	if numCount == 1 {
		n.dp[0] = nums[0]
		return *n
	}

	n.dp[0] = nums[0]

	for i := 1; i < numCount; i++ {
		n.dp[i] = nums[i] + n.dp[i-1]
	}

	return *n
}

func (this *NumArray) SumRange(i int, j int) int {

	if i == 0 {
		return this.dp[j]
	}

	return this.dp[j] - this.dp[i-1]
}

func main() {
	Constructor([]int{})
}
