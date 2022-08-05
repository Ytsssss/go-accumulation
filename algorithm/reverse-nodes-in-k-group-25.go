package algorithm

func ReverseKGroup(head *ListNode, k int) *ListNode {
	firstNode := &ListNode{
		Next: head,
	}
	pre, tail := firstNode, head
	finished := false
	for head != nil {
		index := 0
		for index < k {
			if tail != nil {
				tail = tail.Next
				index++
			} else {
				finished = true
				break
			}
		}
		if finished {
			break
		}
		first, last := reverseNode(head, tail)
		pre.Next = first
		last.Next = tail
		pre = last
		head = tail
	}

	return firstNode.Next
}

func reverseNode(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil && cur != tail {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre, head
}
