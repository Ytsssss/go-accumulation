package algorithm

import (
	"sort"
	"strconv"
)

type SortNum []int

func LargestNumber(nums []int) string {
	sort.Sort(SortNum(nums))
	result := ""
	for len(nums) > 1 {
		if nums[0] == 0 {
			nums = nums[1:]
		} else {
			break
		}
	}

	for _, num := range nums {
		result += strconv.Itoa(num)
	}
	return result
}

func (x SortNum) Len() int { return len(x) }
func (x SortNum) Less(i, j int) bool {
	depthi, depthj, numi, numj := 1, 1, x[i], x[j]
	for numi/10 > 0 {
		numi = numi / 10
		depthi++
	}
	for numj/10 > 0 {
		numj = numj / 10
		depthj++
	}
	curNumi, curNumj := x[i], x[j]
	for k := 0; k < depthi; k++ {
		curNumj *= 10
	}
	for m := 0; m < depthj; m++ {
		curNumi *= 10
	}
	tmpi := curNumi + x[j]
	tmpj := curNumj + x[i]

	return tmpi > tmpj
}
func (x SortNum) Swap(i, j int) { x[i], x[j] = x[j], x[i] }
