package algorithm

type MyStack struct {
	queue []int
}

func Constructor2() MyStack {
	return MyStack{
		queue: []int{},
	}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	tmp := make([]int, 0, len(this.queue))
	for len(this.queue) > 1 {
		tmp = append(tmp, this.queue[0])
		this.queue = this.queue[1:]
	}
	answer := this.queue[0]
	this.queue = tmp

	return answer
}

func (this *MyStack) Top() int {
	tmp := make([]int, 0, len(this.queue))
	for len(this.queue) > 1 {
		tmp = append(tmp, this.queue[0])
		this.queue = this.queue[1:]
	}
	answer := this.queue[0]
	tmp = append(tmp, answer)
	this.queue = tmp

	return answer
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
