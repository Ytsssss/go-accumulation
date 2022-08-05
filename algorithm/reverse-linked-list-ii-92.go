package algorithm

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	tmp := head
	num := 1
	for num < left-1 {
		head = head.Next
		num++
	}
	pre, curTmp := head, head.Next
	cur, nex := reverse(head.Next, right-left)
	pre.Next = cur
	curTmp.Next = nex
	return tmp
}

func reverse(head *ListNode, count int) (*ListNode, *ListNode) {
	var pre *ListNode
	for count >= 0 && head != nil {
		nex := head.Next
		head.Next = pre
		pre = head
		head = nex
		count--
	}
	return pre, head
}
