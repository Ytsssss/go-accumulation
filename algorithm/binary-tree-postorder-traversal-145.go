package algorithm

func postorderTraversal(root *TreeNode) (nodes []int) {
	return postorder(nodes, root)
}

func postorder(nodes []int, node *TreeNode) []int {
	// 后序遍历，按 左子树-右子树-根节点 的顺序输出
	if node == nil {
		return nodes
	}
	nodes = postorder(nodes, node.Left)
	nodes = postorder(nodes, node.Right)
	nodes = append(nodes, node.Val)
	return nodes
}
