package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) (res [][]int) {

	if root == nil{
		return res
	}

	res = make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		l := len(queue)
		ans := make([]int, 0)
		for i := 0; i < l; i++ {

			ans = append(ans, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}

		}

		queue = queue[l:]
		res = append(res, ans)
	}

	r := len(res) - 1
	l := 0

	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}

	return res
}

func main() {

}
