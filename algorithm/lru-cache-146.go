package algorithm

type LRUCache struct {
	length, cap int
	indexMap    map[int]*nodeQueue
	head, end   *nodeQueue
}

type nodeQueue struct {
	// 保存前序和后续节点，构建双向链表
	pre, next  *nodeQueue
	key, value int
}

func ConstructorLRU(capacity int) LRUCache {
	head, end := &nodeQueue{}, &nodeQueue{}
	head.next = end
	end.pre = head
	return LRUCache{
		length:   0,
		cap:      capacity,
		indexMap: make(map[int]*nodeQueue),
		head:     head,
		end:      end,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.indexMap[key]; ok {
		answer := node.value
		node.removeNode()
		node.addNode(this.head)
		return answer
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.indexMap[key]; ok {
		node.removeNode()
		newNode := &nodeQueue{
			key:   key,
			value: value,
		}
		newNode.addNode(this.head)
		this.indexMap[key] = newNode
	} else {
		if this.length >= this.cap {
			delete(this.indexMap, this.end.pre.key)
			this.end.pre.removeEnd(this.end)
		}
		newNode := &nodeQueue{
			key:   key,
			value: value,
		}
		newNode.addNode(this.head)
		this.length++
		this.indexMap[key] = newNode
	}
}

func (n *nodeQueue) removeNode() {
	n.pre.next = n.next
	n.next.pre = n.pre
}

func (n *nodeQueue) removeEnd(end *nodeQueue) {
	n.pre.next = end
	end.pre = n.pre
}

func (n *nodeQueue) addNode(pre *nodeQueue) {
	tmp := pre.next
	pre.next = n
	n.pre = pre
	n.next = tmp
	tmp.pre = n
}
