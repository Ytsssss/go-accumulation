package utils

// DiffStringSlice 判断两个切片的差异，两个返回值为 newElems-oldElems， oldElems-newElems
func DiffStringSlice(newElems []string, oldElems []string) ([]string, []string) {
	strMap := make(map[string]struct{})
	for _, elem := range newElems {
		strMap[elem] = struct{}{}
	}
	var add []string
	var del []string
	for _, elem := range oldElems {
		if _, ok := strMap[elem]; !ok {
			del = append(del, elem)
		}
		delete(strMap, elem)
	}

	if len(strMap) > 0 {
		for k := range strMap {
			add = append(add, k)
		}
	}

	return add, del
}

func Sum[T int | float64](a, b T) T {
	return a + b
}
