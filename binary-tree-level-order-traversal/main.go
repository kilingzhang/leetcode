package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans [][]int

func levelOrder(root *TreeNode) [][]int {

	ans = [][]int{}

	if root == nil {
		return ans
	}

	find(root, 0)
	return ans
}

func find(root *TreeNode, d int) {

	if root == nil {
		return
	}

	if len(ans) == d {
		ans = append(ans, []int{})
	}

	ans[d] = append(ans[d], root.Val)

	if root.Left != nil {
		find(root.Left, d+1)
	}

	if root.Right != nil {
		find(root.Right, d+1)
	}
}


func main() {

}
