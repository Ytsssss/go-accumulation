package algorithm

import "sort"

func WiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	maxLength := 1
	if nums[0] != nums[1] {
		maxLength = 2
	}
	large := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		if (large >= 0 && nums[i] < nums[i-1]) || (large <= 0 && nums[i] > nums[i-1]) {
			large = nums[i] - nums[i-1]
			maxLength++
		}
	}

	return maxLength
}

func WiggleSort1(nums []int) {
	sort.Ints(nums)
	i := (len(nums) + 1) / 2
	for i < (len(nums)+1)/2 {
		k := i
		j := (len(nums) + 1) / 2
		for k < (len(nums)+1)/2 {
			swap(nums, k, j)
			k++
			j++
		}
		i++
	}
}

func WiggleSort(nums []int) {
	sort.Ints(nums)
	nums1 := make([]int, 0)
	nums2 := make([]int, 0)
	nums1 = append(nums1, nums[:(len(nums)+1)/2]...)
	nums2 = append(nums2, nums[(len(nums)+1)/2:]...)

	i, j, k := len(nums1)-1, len(nums2)-1, 0
	for i >= 0 && j >= 0 {
		nums[k] = nums1[i]
		if k+1 < len(nums) {
			nums[k+1] = nums2[j]
		}
		i--
		j--
		k += 2
	}
	if i >= 0 {
		nums[k] = nums1[i]
	}
}

func swap(nums []int, i int, i2 int) {
	tmp := nums[i]
	nums[i] = nums[i2]
	nums[i2] = tmp
}
