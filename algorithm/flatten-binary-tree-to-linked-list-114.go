package algorithm

func Flatten(root *TreeNode) {
	for root != nil {
		left := root.Left
		right := root.Right
		if left == nil {
			root = root.Right
			continue
		}
		root.Right = left
		root.Left = nil
		for left != nil && left.Right != nil {
			left = left.Right
		}
		if left != nil {
			left.Right = right
		}
		root = root.Right
	}
}
