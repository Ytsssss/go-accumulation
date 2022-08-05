package algorithm

func MergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	return merge(root1, root2, &TreeNode{})
}

func merge(root1, root2, new *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return new
	}
	if new == nil {
		new = &TreeNode{}
	}
	if root1 != nil && root2 != nil {
		new.Val = root1.Val + root2.Val
		new.Left = merge(root1.Left, root2.Left, new.Left)
		new.Right = merge(root1.Right, root2.Right, new.Right)
	} else if root1 != nil {
		new.Val = root1.Val
		new.Left = merge(root1.Left, root2, new.Left)
		new.Right = merge(root1.Right, root2, new.Right)
	} else if root2 != nil {
		new.Val = root2.Val
		new.Left = merge(root1, root2.Left, new.Left)
		new.Right = merge(root1, root2.Right, new.Right)
	}
	return new
}
