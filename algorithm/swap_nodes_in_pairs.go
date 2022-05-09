package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	headNode := &ListNode{
		Val:  0,
		Next: head,
	}
	pre := headNode
	for pre.Next != nil && pre.Next.Next != nil {
		cur := pre.Next
		next := cur.Next
		suf := next.Next
		pre.Next = next
		next.Next = cur
		cur.Next = suf
		pre = cur
	}
	return headNode.Next
}
