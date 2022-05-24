package algorithm

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	nums := make([]int, 0)
	return appendNode(root, nums)
}

func appendNode(node *TreeNode, nums []int) []int {
	if node == nil {
		return nums
	}
	nums = appendNode(node.Left, nums)
	nums = append(nums, node.Val)
	nums = appendNode(node.Right, nums)
	return nums
}

func inorderTraversal2(root *TreeNode) []int {
	nums := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nums = append(nums, root.Val)
		root = root.Right
	}

	return nums
}
