package main

import (
	"fmt"
)

func rotate(matrix [][]int) {

	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {

		n := 0
		m := len(matrix[i]) - 1
		for n < m {
			matrix[i][n], matrix[i][m] = matrix[i][m], matrix[i][n]
			n++
			m--
		}

	}

}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Println(matrix)
}
