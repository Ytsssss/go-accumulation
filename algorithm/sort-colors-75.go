package algorithm

func SortColors(nums []int) {
	p, q, i := 0, len(nums)-1, 0
	for i <= q {
		if nums[i] == 0 {
			nums[i], nums[p] = nums[p], nums[i]
			p++
			for nums[p] == 0 {
				p++
			}
		}
		if nums[i] == 2 {
			nums[i], nums[q] = nums[q], nums[i]
			q--
			for nums[q] == 2 {
				q--
			}
		}
		i++
	}
}
