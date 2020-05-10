package main

import (
	"fmt"
	tool "github.com/kilingzhang/leetcode"
)

func spiralOrder(matrix [][]int) (res []int) {

	h := len(matrix) - 1

	if h == -1 {
		return res
	}

	w := len(matrix[0]) - 1

	m := 0
	n := 0

	visited := make([][]int, h+1)

	for i := 0; i <= h; i++ {
		visited[i] = make([]int, w+1)
	}

	for h-m >= 0 && w-n >= 0 {
		fmt.Println(m, n)
		res = append(res, right(matrix, visited, m, n)...)
		fmt.Println(m, w-n)
		res = append(res, down(matrix, visited, m, w-n)...)
		fmt.Println(h-m, w-n)
		res = append(res, left(matrix, visited, h-m, w-n)...)
		fmt.Println(h-m, n)
		res = append(res, up(matrix, visited, h-m, n)...)
		fmt.Println(res, visited)
		m++
		n++
	}

	return res
}

func right(matrix [][]int, visited [][]int, m int, n int) (res []int) {

	for i := n; i < len(matrix[m]); i++ {
		if visited[m][i] == 0 {
			res = append(res, matrix[m][i])
			visited[m][i] = 1
		}
	}
	return res
}

func down(matrix [][]int, visited [][]int, m int, n int) (res []int) {

	for i := 0; i < len(matrix); i++ {
		if visited[i][n] == 0 {
			res = append(res, matrix[i][n])
			visited[i][n] = 1
		}
	}
	return res
}

func left(matrix [][]int, visited [][]int, m int, n int) (res []int) {

	for i := n; i >= 0; i-- {
		if visited[m][i] == 0 {
			res = append(res, matrix[m][i])
			visited[m][i] = 1
		}
	}
	return res
}

func up(matrix [][]int, visited [][]int, m int, n int) (res []int) {

	for i := m; i >= 0; i-- {
		if visited[i][n] == 0 {
			res = append(res, matrix[i][n])
			visited[i][n] = 1
		}
	}
	return res
}

func main() {
	tool.Dump(spiralOrder([][]int{{1}}))
}
