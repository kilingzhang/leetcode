package main

import (
	"fmt"
)

func generate(numRows int) [][]int {

	if numRows == 1 {
		return [][]int{{1}}
	}

	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}

	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][i] = 1
		res[i][0] = 1
	}

	for i := 2; i < numRows; i++ {
		for j := 1; j < i; j++ {
			fmt.Println(i, j, res)
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}

	return res
}

func main() {
	fmt.Println(generate(5))
}
