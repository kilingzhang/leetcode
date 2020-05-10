package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) (tree *TreeNode) {
	parent := map[int]*TreeNode{}
	visited := map[int]bool{}

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {

		if root == nil {
			return
		}

		if root.Left != nil {
			parent[root.Left.Val] = root
			dfs(root.Left)
		}

		if root.Right != nil {
			parent[root.Right.Val] = root
			dfs(root.Right)
		}

	}

	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}

	return nil
}


func main() {

}
