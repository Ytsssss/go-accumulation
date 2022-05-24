package algorithm

type MyQueue struct {
	// 建立两个 stack 一个用于存储字符串，一个用于查询和出栈，模拟队列功能
	storeStack []int
	queryStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		storeStack: []int{},
		queryStack: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.storeStack = append(this.storeStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.queryStack) == 0 {
		this.exchange()
	}
	i := this.queryStack[len(this.queryStack)-1]
	this.queryStack = this.queryStack[:len(this.queryStack)-1]
	return i
}

func (this *MyQueue) exchange() {
	for len(this.storeStack) > 0 {
		this.queryStack = append(this.queryStack, this.storeStack[len(this.storeStack)-1])
		this.storeStack = this.storeStack[:len(this.storeStack)-1]
	}
}

func (this *MyQueue) Peek() int {
	if len(this.queryStack) == 0 {
		this.exchange()
	}
	return this.queryStack[len(this.queryStack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.storeStack) == 0 && len(this.queryStack) == 0
}
