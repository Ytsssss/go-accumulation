package algorithm

//请你设计并实现一个满足 LRU (最近最少使用) 缓存 约束的数据结构。
//实现 LRUCache 类：
//LRUCache(int capacity) 以 正整数 作为容量capacity 初始化 LRU 缓存
//int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
//void put(int key, int value)如果关键字key 已经存在，则变更其数据值value ；如果不存在，则向缓存中插入该组key-value 。如果插入操作导致关键字数量超过capacity ，则应该 逐出 最久未使用的关键字。
//函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

//解释
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // 缓存是 {1=1}
//lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
//lRUCache.get(1);    // 返回 1
//lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
//lRUCache.get(2);    // 返回 -1 (未找到)
//lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
//lRUCache.get(1);    // 返回 -1 (未找到)
//lRUCache.get(3);    // 返回 3
//lRUCache.get(4);    // 返回 4

type LRUCache0 struct {
	Head, Tail *Node
	Cap        int
	Length     int
	NodeMap    map[int]*Node
}

type Node struct {
	Val       int
	Pre, Next *Node
}

func Constructor0(capacity int) LRUCache0 {
	empty := &Node{
		Val: -1,
	}
	head, tail := empty, empty
	head.Next = tail
	tail.Pre = head
	return LRUCache0{
		Head:    head,
		Tail:    tail,
		Length:  0,
		Cap:     capacity,
		NodeMap: make(map[int]*Node),
	}
}

func (this *LRUCache0) Get(key int) int {
	v, ok := this.NodeMap[key]
	if !ok {
		return -1
	}
	res := v.Val
	this.putAfterHead(v)
	return res
}

func (this *LRUCache0) Put(key int, value int) {
	if v, ok := this.NodeMap[key]; ok {
		v.Val = value
		this.putAfterHead(v)
	} else {
		newNode := &Node{
			Val: value,
		}
		this.putHeadNext(newNode)

		if this.Length+1 > this.Cap {
			this.removeTail()
		} else {
			this.Length++
		}
	}
}

func (this *LRUCache0) putAfterHead(v *Node) {
	pre, next := v.Pre, v.Next
	pre.Next = next
	next.Pre = pre

	this.putHeadNext(v)
}

func (this *LRUCache0) putHeadNext(v *Node) {
	head := this.Head
	hnext := head.Next
	head.Next = v
	v.Pre = head
	v.Next = hnext
	hnext.Pre = v
}

func (this *LRUCache0) removeTail() {
	tail := this.Tail
	pre := tail.Pre.Pre
	pre.Next = tail
	tail.Pre = pre
}
