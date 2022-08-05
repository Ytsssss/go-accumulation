package algorithm

func minDepth(root *TreeNode) int {
	return minDepthCount(root, 0)
}

func minDepthCount(root *TreeNode, count int) int {
	if root == nil {
		return count
	}
	if root.Left == nil {
		return minDepthCount(root.Right, count+1)
	}
	if root.Right == nil {
		return minDepthCount(root.Left, count+1)
	}
	left := minDepthCount(root.Left, count+1)
	right := minDepthCount(root.Right, count+1)
	if left > right {
		return right
	} else {
		return left
	}
}

func minDepthDeepLoop(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	count := 0
	for len(queue) > 0 {
		count++
		current := make([]*TreeNode, 0)
		for i := 0; i < len(queue); i++ {
			if queue[i].Left == nil && queue[i].Right == nil {
				return count
			}
			if queue[i].Left == nil {
				current = append(current, queue[i].Right)
			} else if queue[i].Right == nil {
				current = append(current, queue[i].Left)
			} else {
				current = append(current, queue[i].Left, queue[i].Right)
			}
		}
		queue = current
	}
	return count
}
