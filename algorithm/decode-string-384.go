package algorithm

import (
	"bytes"
	"strconv"
)

type numString struct {
	Repeat int
	Str    string
}

func DecodeString(s string) string {
	buffer := bytes.Buffer{}
	numInt, tmpStr, hasLeft := 0, "", false
	stack := make([]*numString, 0)
	for i := 0; i < len(s); i++ {
		if num, ok := isNumber(s[i]); ok {
			numInt = numInt*10 + num
		} else if s[i] == '[' {
			stack = append(stack, &numString{
				Repeat: numInt,
				Str:    "",
			})
			if !hasLeft {
				buffer.WriteString(tmpStr)
				tmpStr = ""
			}
			numInt = 0
			hasLeft = true
		} else if s[i] == ']' {
			tmp := stack[len(stack)-1]
			tmpBuffer := bytes.Buffer{}
			for tmp.Repeat > 0 {
				tmpBuffer.WriteString(tmp.Str)
				tmp.Repeat--
			}
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				stack[len(stack)-1].Str = stack[len(stack)-1].Str + tmpBuffer.String()
			} else {
				buffer.WriteString(tmpBuffer.String())
				hasLeft = false
			}
		} else if !hasLeft {
			tmpStr = tmpStr + string(s[i])
		} else {
			stack[len(stack)-1].Str = stack[len(stack)-1].Str + string(s[i])
		}
	}
	buffer.WriteString(tmpStr)
	return buffer.String()
}

func isNumber(c byte) (int, bool) {
	num, err := strconv.Atoi(string(c))
	if err != nil {
		return 0, false
	}
	return num, true
}
