package main

import (
	"fmt"
)

func main() {

	var grid [][] int = [][] int{
		{
			2, 2, 2,
		},
		{
			0, 1, 0,
		},
		{
			1, 0, 2,
		},
	}

	n := orangesRotting(grid)

	fmt.Println(n)

}

func isNewOranges(grids [][]int) bool {
	for _, grid := range grids {
		for _, v := range grid {
			if v == 1 {
				return true
			}
		}
	}

	return false
}

func isContinueDiffuseDecay(grids [][]int) (isContinue bool) {

	isContinue = false

	for i, grid := range grids {
		for j, v := range grid {

			if v == 2 {

				if i-1 >= 0 && grids[i-1][j] == 1 {
					isContinue = true
				}

				if i+1 < len(grids) && grids[i+1][j] == 1 {
					isContinue = true
				}

				if j-1 >= 0 && grids[i][j-1] == 1 {
					isContinue = true
				}

				if j+1 < len(grid) && grids[i][j+1] == 1 {
					isContinue = true
				}

			}

		}
	}

	return isContinue
}

func diffuseDecay(grids [][]int) [][]int {

	for i, grid := range grids {
		for j, v := range grid {

			if v == 2 {

				if i-1 >= 0 && grids[i-1][j] == 1 {
					grids[i-1][j] = 3
				}

				if i+1 < len(grids) && grids[i+1][j] == 1 {
					grids[i+1][j] = 3
				}

				if j-1 >= 0 && grids[i][j-1] == 1 {
					grids[i][j-1] = 3
				}

				if j+1 < len(grid) && grids[i][j+1] == 1 {
					grids[i][j+1] = 3
				}

			}

		}
	}

	for i, grid := range grids {
		for j, v := range grid {
			if v == 3 {
				grids[i][j] = 2
			}
		}
	}

	return grids
}

func orangesRotting(grids [][]int) int {

	var n int = 0

	for {

		if !isNewOranges(grids) {

			return n
		}

		if !isContinueDiffuseDecay(grids) {
			if !isNewOranges(grids) {
				return n
			}
			return -1
		}

		grids = diffuseDecay(grids)

		n++
	}

}
