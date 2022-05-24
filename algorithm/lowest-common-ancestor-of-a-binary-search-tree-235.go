package algorithm

// 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 二叉搜索树有序，左子树必小于当前节点，右子树必大于当前节点
	if root == nil || root == p || root == q {
		// 如果当前节点为空，或者当前节点和p或q中某一个相等，则当前节点为p和q的最近公共祖先
		return root
	}
	if root.Val > p.Val && root.Val > q.Val {
		// 如果当前节点的值比p和q都大，则表示最近公共祖先在左子树
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		// 如果当前节点的值比p和q都小，则表示最近公共祖先在右子树
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		// 如果p和q一个比当前节点值大，一个比当前节点值小，则表示最近公共祖先就为当前节点
		return root
	}
}
