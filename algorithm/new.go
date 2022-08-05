package algorithm

// 给个 list 输出所有可能的 subsets
//给出一个长度n的list，输出list所有可能的子集合
//例如：
//给出的list为[1,2,3]
//输出为[1],[2],[3],[1,2],[1,3],[2,3],[1,2,3]
// [1,2,3]
// [1,2,3]
// [1,2,3]
var result = [][]int{}

func GetList(list []int) [][]int {
	getAnswer2(list, make([]int, 0), 0)
	return result
}

func getAnswer2(list []int, valueMap []int, depth int) {
	if depth == len(list) {
		result = append(result, valueMap)
		return
	}
	tmp := append(valueMap, list[depth])
	getAnswer2(list, tmp, depth+1)
	getAnswer2(list, valueMap, depth+1)

}
