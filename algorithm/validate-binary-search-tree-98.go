package algorithm

import "math"

func isValidBST(root *TreeNode) bool {
	// 中序遍历二叉树，如果所得数组有序，表示该二叉树为二叉搜索树
	var nodes []int
	nodes = appendNode(root, nodes)
	for i := 1; i <= len(nodes); i++ {
		if nodes[i] < nodes[i-1] {
			return false
		}
	}
	return true
}

func isValidBST2(root *TreeNode) bool {
	// 中序遍历二叉树，每次记录前一个数，将当前节点和前一个节点进行比较，如果当前节点不大于前一个节点则不为二叉搜索树
	stack := make([]*TreeNode, 0)
	preNum := math.MinInt
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= preNum {
			return false
		}
		preNum = root.Val
		root = root.Right
	}
	return true
}

func isValidBST3(root *TreeNode) bool {
	// 二叉搜索树根节点一定比左子树大，一定比右子树小
	return checkNode(root, math.MinInt, math.MaxInt)
}

func checkNode(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	return checkNode(node.Left, min, node.Val) && checkNode(node.Right, node.Val, max)
}
