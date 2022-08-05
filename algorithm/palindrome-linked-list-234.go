package algorithm

func IsPalindrome(head *ListNode) bool {

	return true
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

func endOfFirstHalf(head *ListNode) *ListNode {
	//fast, slow := head, head
	return nil

}
