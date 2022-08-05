package algorithm

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	flow, tmp := 0, l1
	for l1 != nil && l1.Next != nil {
		if l2 != nil {
			num := l1.Val + l2.Val + flow
			if num >= 10 {
				num = num - 10
				flow = 1
			} else {
				flow = 0
			}
			l1.Val = num
			l1 = l1.Next
			l2 = l2.Next
		} else {
			num := l1.Val + flow
			if num >= 10 {
				num = num - 10
				flow = 1
			} else {
				flow = 0
			}
			l1.Val = num
			l1 = l1.Next
		}
	}
	if l1 == nil {
		l1 = l2
	} else if l1.Next == nil {
		l1.Next = l2
	}
	for l1 != nil {
		num := l1.Val + flow
		if num >= 10 {
			num = num - 10
			flow = 1
		} else {
			flow = 0
		}
		l1.Val = num
		l1 = l1.Next
	}
	return tmp
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	newNums := make([]int, 0, len(nums1)+len(nums2))
	left, right := 0, 0
	for left < len(nums1) && right < len(nums2) {
		if nums1[left] < nums2[right] {
			newNums = append(newNums, nums1[left])
			left++
		} else {
			newNums = append(newNums, nums2[right])
			right++
		}
	}
	if left < len(nums1) {
		newNums = append(newNums, nums1[left:]...)
	}
	if right < len(nums2) {
		newNums = append(newNums, nums2[right:]...)
	}
	if (len(nums1)+len(nums2))%2 != 0 {
		return float64(newNums[len(newNums)/2])
	} else {
		return float64(newNums[len(newNums)/2]) + float64(newNums[len(newNums)/2-1])/2
	}
}
