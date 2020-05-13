package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) (deep int) {
	if root == nil {
		return deep
	}
	return _minDepth(root, 0)
}

func _minDepth(root *TreeNode, lastDeep int) (deep int) {

	if root == nil {
		return lastDeep
	}

	deep = lastDeep + 1
	ld := deep
	rd := deep

	if root.Left != nil {
		ld = _minDepth(root.Left, deep)
	}

	if root.Right != nil {
		rd = _minDepth(root.Right, deep)
	}

	if root.Left != nil && root.Right == nil {
		return ld
	}

	if root.Right != nil && root.Left == nil {
		return rd
	}

	if root.Left == nil && root.Right == nil {
		return deep
	}

	return int(math.Min(float64(ld), float64(rd)))
}

func main() {

}
