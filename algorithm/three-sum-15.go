package algorithm

import "sort"

func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	m := make(map[int]int, 0)
	for i, num := range nums {
		m[-num] = i
	}
	answerMap := make(map[[3]int]bool, 0)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if k, ok := m[nums[i]+nums[j]]; ok && k != i && k != j {
				ints := [3]int{nums[i], nums[j], nums[k]}
				sort.Ints(ints[:])
				answerMap[ints] = true
			}
		}
	}
	answer := make([][]int, 0)
	for ints := range answerMap {
		elems := make([]int, len(ints))
		copy(elems, ints[:])
		answer = append(answer, elems)
	}

	return answer
}

func ThreeSum2(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	var answer [][]int
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		// 此处判断是否和前一个数相同，是为了避免出现重复结果
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if nums[i]+nums[left]+nums[right] == 0 {
				answer = append(answer, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return answer
}
