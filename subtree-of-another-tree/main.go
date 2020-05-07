package main

import tool "github.com/kilingzhang/leetcode"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return checkTree(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func checkTree(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val == b.Val {
		return checkTree(a.Left, b.Left) && checkTree(a.Right, b.Right)
	}
	return false
}

func main() {
	tool.Dump()
}
