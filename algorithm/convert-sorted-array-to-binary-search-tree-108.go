package algorithm

func SortedArrayToBST(nums []int) *TreeNode {
	node := &TreeNode{}
	if len(nums) == 0 {
		return nil
	}
	left, right := 0, len(nums)-1
	middle := (left + right) / 2
	node.Val = nums[middle]
	node.Left = SortedArrayToBST(nums[:middle])
	node.Right = SortedArrayToBST(nums[middle+1:])

	return node
}

func BalanceBST(root *TreeNode) *TreeNode {
	nums := middle(root, make([]int, 0))
	tree := buildTree(nums, 0, len(nums)-1)
	return tree
}

func buildTree(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	node := &TreeNode{
		Val: nums[mid],
	}
	node.Left = buildTree(nums, left, mid-1)
	node.Right = buildTree(nums, mid+1, right)
	return node
}

func middle(root *TreeNode, nums []int) []int {
	if root == nil {
		return nums
	}
	nums = middle(root.Left, nums)
	nums = append(nums, root.Val)
	nums = middle(root.Right, nums)
	return nums
}
