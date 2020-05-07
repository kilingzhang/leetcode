package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validBST(root, math.MinInt64, math.MaxInt64)
}

func validBST(root *TreeNode, min int, max int) bool {

	if root == nil {
		return true
	}

	if root.Right == nil && root.Left == nil {
		return true
	}

	if root.Left != nil && (root.Val <= root.Left.Val || root.Left.Val <= min) {
		return false
	}

	if root.Right != nil && (root.Val >= root.Right.Val || root.Right.Val >= max) {
		return false
	}

	if root.Left != nil && !validBST(root.Left, min, root.Val) {
		return false
	}

	if root.Right != nil && !validBST(root.Right, root.Val, max) {
		return false
	}

	return true
}

func main() {

}
