package algorithm

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	if root == nil {
		return result
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		currentResult := make([]int, 0)
		currentQueue := make([]*TreeNode, 0)
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			currentResult = append(currentResult, node.Val)
			if node.Left != nil {
				currentQueue = append(currentQueue, node.Left)
			}
			if node.Right != nil {
				currentQueue = append(currentQueue, node.Right)
			}
		}
		queue = currentQueue
		result = append(result, currentResult)
	}

	return result
}

func levelOrderDeep(root *TreeNode) [][]int {
	result := make([][]int, 0)
	index := 0
	result = deepLevelOrder(root, result, index)
	return result
}

func deepLevelOrder(root *TreeNode, result [][]int, index int) [][]int {
	if root == nil {
		return result
	}
	if index > len(result)-1 {
		result = append(result, []int{})
	}
	result[index] = append(result[index], root.Val)
	result = deepLevelOrder(root.Left, result, index+1)
	result = deepLevelOrder(root.Right, result, index+1)
	return result
}
