package algorithm

func permute(nums []int) [][]int {
	return getResult(nums, []int{}, [][]int{}, 0, map[int]bool{})
}

func getResult(nums, current []int, answer [][]int, depth int, lastIndex map[int]bool) [][]int {
	if depth == len(nums) {
		answer = append(answer, current)
		return answer
	}
	for i := 0; i < len(nums); i++ {
		if _, ok := lastIndex[i]; !ok {
			lastIndex[i] = true
			tmp := make([]int, 0)
			tmp = current[:]
			current = append(current, nums[i])
			answer = getResult(nums, current, answer, depth+1, lastIndex)
			delete(lastIndex, i)
			current = tmp
		}
	}
	return answer
}
