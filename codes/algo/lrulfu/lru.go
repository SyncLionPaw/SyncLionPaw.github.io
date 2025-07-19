package main

type Node struct {
	// value ...
	prev  *Node
	next  *Node
	value int
	key   int
}

type DoubleLinkList struct {
	head *Node
	tail *Node
	size int
}

func NewDoubleLinkList() *DoubleLinkList {
	h, t := &Node{prev: nil, next: nil}, &Node{prev: nil, next: nil}
	h.next = t
	t.prev = h
	return &DoubleLinkList{h, t, 0}
}

// pick up node from list
func (d *DoubleLinkList) Pick(node *Node) {
	// assert node in list and node isn't head and tail
	node.prev.next = node.next
	node.next.prev = node.prev
	d.size--
}

func (d *DoubleLinkList) InsertHead(node *Node) {
	if node == nil {
		return
	}
	hn := d.head.next
	d.head.next = node
	node.prev = d.head

	node.next = hn
	hn.prev = node
	d.size++
}

func (d *DoubleLinkList) InsertTail(node *Node) {
	if node == nil {
		return
	}
	tp := d.tail.prev
	d.tail.prev = node
	node.next = d.tail

	node.prev = tp
	tp.next = node
	d.size++
}

type LRUCache struct {
	cap int
	size  int
	link  *DoubleLinkList
	knMap map[int]*Node
}

func Constructor(capacity int) LRUCache {
    return LRUCache{capacity, 0, NewDoubleLinkList(), make(map[int]*Node)}
}


func (this *LRUCache) Get(key int) int {
    node, e := this.knMap[key]
	if !e {
		return -1
	}
	this.link.Pick(node)
	this.link.InsertHead(node)
	return node.value
}


func (this *LRUCache) Put(key int, value int)  {
    node, e := this.knMap[key]
	if e {
		node.value = value
		this.link.Pick(node)
		this.link.InsertHead(node)
		return
	}
	newNode := &Node{key: key, value: value}
	if this.cap > this.size {
		this.link.InsertHead(newNode)
		this.link.size++
	} else {
		lastOne := this.link.tail.prev
		this.link.Pick(lastOne)
		delete(this.knMap, lastOne.key)
		this.link.InsertHead(newNode)
	}
	this.knMap[key] = newNode
	return
}
