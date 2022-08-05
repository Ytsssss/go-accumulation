package algorithm

import "sort"

func CombinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return dfs1(candidates, make([]int, 0), target, 0, make([][]int, 0))
}

func dfs1(candidates, cur []int, target, index int, answer [][]int) [][]int {
	if target == 0 {
		return append(answer, append([]int(nil), cur...))
	}
	for i := index; i < len(candidates); i++ {
		if target-candidates[i] < 0 {
			break
		}
		answer = dfs1(candidates, append(cur, candidates[i]), target-candidates[i], i, answer)
	}
	return answer
}
