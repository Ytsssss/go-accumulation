package algorithm

import (
	"container/heap"
	"sort"
)

type KthLargest struct {
	k int
	// 用于存储最大的K个数，模拟构建一个小顶堆
	topNums []int
}

func Constructor3(k int, nums []int) KthLargest {
	kthLargest := KthLargest{
		k:       k,
		topNums: make([]int, 0, k),
	}
	for _, num := range nums {
		kthLargest.Add(num)
	}
	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	if len(this.topNums) < this.k {
		this.push(val)
	} else if val > this.top() {
		this.pop()
		this.push(val)
	}

	return this.top()
}

func (this *KthLargest) push(val int) {
	left, right, tmp := 0, len(this.topNums), 0
	if left == right {
		this.topNums = append(this.topNums, val)
		return
	}
	for left < right {
		tmp = (left + right) / 2
		if this.topNums[tmp] >= val {
			tmp++
			left = tmp
		} else {
			right = tmp
		}
	}
	nums := make([]int, 0, len(this.topNums)+1)
	nums = append(nums, this.topNums[:tmp]...)
	nums = append(nums, val)
	nums = append(nums, this.topNums[tmp:]...)
	this.topNums = nums
}

func (this *KthLargest) pop() {
	this.topNums = this.topNums[:len(this.topNums)-1]
}

func (this *KthLargest) top() int {
	return this.topNums[len(this.topNums)-1]
}

type KthLargest2 struct {
	k int
	// 用于存储最大的K个数，模拟构建一个小顶堆
	sort.IntSlice
}

func Constructor4(k int, nums []int) KthLargest2 {
	kthLargest := KthLargest2{
		k: k,
	}
	for _, num := range nums {
		kthLargest.Add(num)
	}
	return kthLargest
}

func (this *KthLargest2) Add(val int) int {
	heap.Push(this, val)
	if this.Len() > this.k {
		heap.Pop(this)
	}
	return this.top().(int)
}

func (this *KthLargest2) Push(x interface{}) {
	this.IntSlice = append(this.IntSlice, x.(int))
}

func (this *KthLargest2) Pop() interface{} {
	tmp := this.IntSlice[this.Len()-1]
	this.IntSlice = this.IntSlice[:this.Len()-1]
	return tmp
}

func (this *KthLargest2) top() interface{} {
	return this.IntSlice[0]
}
