package algorithm

import (
	"sort"
	"strings"
)

func diffWaysToCompute(expression string) []int {
	return nil
}

func ReplaceWords(dictionary []string, sentence string) string {
	strSort := &StringSort{
		Strs: dictionary,
	}
	sort.Sort(strSort)
	splits := strings.Split(sentence, " ")
	for i := 0; i < len(splits); i++ {
		for j := 0; j < len(strSort.Strs); j++ {
			if strings.HasPrefix(splits[i], strSort.Strs[j]) {
				splits[i] = strSort.Strs[j]
				break
			}
		}
	}
	result := ""
	for i := 0; i < len(splits); i++ {
		result = result + " " + splits[i]
	}
	return result[1:]
}

type StringSort struct {
	Strs []string
}

func (s *StringSort) Len() int {
	return len(s.Strs)
}

func (s *StringSort) Less(i, j int) bool {
	return len(s.Strs[i]) < len(s.Strs[j])
}

func (s *StringSort) Swap(i, j int) {
	tmp := s.Strs[j]
	s.Strs[j] = s.Strs[i]
	s.Strs[i] = tmp
}
