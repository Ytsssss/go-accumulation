package algorithm

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func hasCycle2(head *ListNode) bool {
	set := make(map[*ListNode]struct{}, 0)
	for head != nil {
		if _, ok := set[head]; ok {
			return true
		}
		set[head] = struct{}{}
		head = head.Next
	}
	return false
}
