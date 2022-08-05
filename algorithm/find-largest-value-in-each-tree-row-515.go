package algorithm

import "math"

func largestValues(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		max := math.MinInt
		tmpQueue := make([]*TreeNode, 0)
		for _, node := range queue {
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				tmpQueue = append(tmpQueue, node.Left)
			}
			if node.Right != nil {
				tmpQueue = append(tmpQueue, node.Right)
			}
		}
		queue = tmpQueue
		result = append(result, max)
	}
	return result
}
