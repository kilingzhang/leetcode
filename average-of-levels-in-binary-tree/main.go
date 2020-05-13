package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) (res []float64) {

	if root == nil {
		return res
	}

	res = make([]float64, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	l := len(queue)
	for l != 0 {

		avg := 0.0
		for i := 0; i < l; i++ {

			avg += float64(queue[i].Val)

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		queue = queue[l:]
		avg /= float64(l)
		res = append(res, avg)
		l = len(queue)
	}

	return res
}

func main() {

}
