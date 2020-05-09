package main

import (
	tool "github.com/kilingzhang/leetcode"
)

func findDiagonalOrder(matrix [][]int) []int {
	// 检查空数组
	if matrix == nil || len(matrix) == 0 {
		return []int{}
	}

	// N行， M列
	n, m := len(matrix), len(matrix[0])

	// 存储结果
	var result []int

	// 对角线处理切片
	var diagonalLine []int

	// 切片清空
	//diagonalLine = diagonalLine[:0]

	// 遍历对角线，第一行和最后一列元素都是对角线的起点。则共有n+m-1条对角线
	for i := 0; i < n+m-1; i++ {
		// 清空对角线切片
		diagonalLine = diagonalLine[:0]

		// 对角线起点横坐标
		var a int
		if i < m {
			a = 0
		} else {
			a = i - m + 1
		}

		// 对角线起点纵坐标
		var o int
		if i < m {
			o = i
		} else {
			o = m - 1
		}

		// 找到对角线起点后，开始遍历当前起点为(a, o)的对角线
		for a < n && o > -1 {
			diagonalLine = append(diagonalLine, matrix[a][o])
			// 右上元素往左下移动，(a+1, j-1)
			a++
			o--
		}

		// 如果为奇数对角线，则当前对角线切片元素取反。(因i从0开始，故对2取模为0时需要反转对角线)
		if i%2 == 0 {
			diagonalLine = Reverse(diagonalLine)
		}

		result = append(result, diagonalLine...)
	}

	return result
}

func Reverse(digits []int) []int {
	i := 0
	j := len(digits) - 1
	for i < j {
		digits[i], digits[j] = digits[j], digits[i]
		i++
		j--
	}
	return digits
}

func main() {
	tool.Dump(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
