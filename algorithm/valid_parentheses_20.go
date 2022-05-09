package algorithm

import "strings"

func IsValid(s string) bool {
	if len(s)%2 != 0 {
		// 如果s长度不为2的倍数，该字符串一定无效
		return false
	}
	validMap := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	stack := make([]byte, 0)
	for _, char := range s {
		if v, ok := validMap[byte(char)]; ok {
			// 如果当前字符为右括号，则需要判断堆栈栈顶的字符是否为对应的左括号
			if len(stack) < 1 || stack[len(stack)-1] != v {
				// 栈顶元素和预期左括号不相等则字符串不合规
				return false
			}
			// 如果匹配成功则移除栈顶元素
			stack = stack[:len(stack)-1]
		} else {
			// 如果当前字符为左括号，将字符入栈
			stack = append(stack, byte(char))
		}
	}

	return len(stack) == 0
}

func IsValid2(s string) bool {
	loop := true
	for loop {
		trim := strings.ReplaceAll(s, "{}", "")
		trim = strings.ReplaceAll(trim, "[]", "")
		trim = strings.ReplaceAll(trim, "()", "")
		loop = len(trim) != len(s)
		s = trim
	}
	return len(s) == 0
}
