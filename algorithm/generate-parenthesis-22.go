package algorithm

func GenerateParenthesis(n int) []string {
	return generate("", []string{}, n, n)
}

func generate(str string, result []string, left, right int) []string {
	if left == 0 && right == 0 {
		result = append(result, str)
		return result
	}
	// 只要剩余的左括号数大于右括号数，就先使用左括号
	if left >= right || left > 0 {
		result = generate(str+"(", result, left-1, right)
	}
	if right > 0 && right > left {
		result = generate(str+")", result, left, right-1)
	}
	return result
}
