package algorithm

// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}
	leftNode := lowestCommonAncestor2(root.Left, p, q)
	rightNode := lowestCommonAncestor2(root.Right, p, q)
	if leftNode == nil {
		// 如果左子树为空，表示左边没有p或q，则公共祖先肯定在右子树
		return rightNode
	}
	if rightNode == nil {
		// 如果右子树为空，表示右边没有p或q，则公共祖先肯定在左子树
		return leftNode
	}
	// 如果左右子树都不为空，则表示p和q分别在左右子树，则公共祖先为当前节点
	return root
}
