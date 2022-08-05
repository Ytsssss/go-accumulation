package algorithm

func preorderTraversal(root *TreeNode) (nodes []int) {
	return preorder(nodes, root)
}

func preorder(nodes []int, node *TreeNode) []int {
	// 前序遍历，按 根节点-左子树-右子树 的顺序输出
	if node == nil {
		return nodes
	}
	nodes = append(nodes, node.Val)
	nodes = preorder(nodes, node.Left)
	nodes = preorder(nodes, node.Right)
	return nodes
}
