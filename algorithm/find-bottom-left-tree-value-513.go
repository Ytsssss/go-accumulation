package algorithm

func findBottomLeftValue(root *TreeNode) int {
	result := root.Val
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		tmpQueue := make([]*TreeNode, 0)
		for i, node := range queue {
			if i == 0 {
				result = node.Val
			}
			if node.Left != nil {
				tmpQueue = append(tmpQueue, node.Left)
			}
			if node.Right != nil {
				tmpQueue = append(tmpQueue, node.Right)
			}
		}
		queue = tmpQueue
	}
	return result
}
