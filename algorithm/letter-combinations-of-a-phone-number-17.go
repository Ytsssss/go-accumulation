package algorithm

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	strMap := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	cur := make([][]byte, 0)
	for i := 0; i < len(digits); i++ {
		cur = append(cur, strMap[digits[i]])
	}
	return getAnswer(cur, 0, []string{}, "")
}

func getAnswer(cur [][]byte, depth int, result []string, str string) []string {
	if depth == len(cur) {
		result = append(result, str)
		return result
	}
	for k := 0; k < len(cur[depth]); k++ {
		result = getAnswer(cur, depth+1, result, str+string(cur[depth][k]))
	}
	return result
}
