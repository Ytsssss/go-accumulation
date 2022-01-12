package algorithm

import "fmt"

// GetStringAllAloud 输入一个字符串，打印出该字符串中字符的所有排列。
func GetStringAllAloud(str string) []string {
	var result []string

	for _, char := range str {
		fmt.Println(string(char))
	}
	return result
}
