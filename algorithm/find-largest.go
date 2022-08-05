package algorithm

import "sort"

func GetLargest10(inputs []string) []string {
	countMap := make(map[string]int)
	for _, input := range inputs {
		countMap[input]++
	}
	list := make([]string, 0)
	for k := range countMap {
		list = append(list, k)
	}
	sort.Slice(list, func(i, j int) bool {
		return countMap[list[i]] < countMap[list[j]]
	})
	return list[:10]
}
