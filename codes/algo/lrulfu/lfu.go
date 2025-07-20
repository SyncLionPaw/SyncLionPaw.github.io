package main

// 淘汰的时候，对于次数一样的，优先去掉最久没用的（lru）

type Node struct {
	// value ...
	prev  *Node
	next  *Node
	value int
	key   int
	freq  int
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

type LFUCache struct {
	cap         int
	size        int
	freqLinkMap map[int]*DoubleLinkList // same freq in same linklist
	knMap       map[int]*Node           // key -> Node
	minFreq     int
}

func Constructor(capacity int) LFUCache {
	flm := make(map[int]*DoubleLinkList)
	knm := make(map[int]*Node)
	return LFUCache{capacity, 0, flm, knm, 1}
}

func (this *LFUCache) GetListByFreq(freq int) *DoubleLinkList {
	l := this.freqLinkMap[freq]
	if l != nil {
		return l
	}
	l = NewDoubleLinkList()
	this.freqLinkMap[freq] = l
	return l
}

func (this *LFUCache) Get(key int) int {
	node, e := this.knMap[key]
	if !e {
		return -1
	}
	// 从原先的频次表上取下，加入到高频的头部
	oldLink := this.freqLinkMap[node.freq]
	oldLink.Pick(node)
	if node.freq == this.minFreq && oldLink.size == 0 { // 旧链已经空了
		this.minFreq++
	}
	node.freq++
	// 加入新的链
	this.GetListByFreq(node.freq).InsertHead(node)
	return node.value
}

func (this *LFUCache) Put(key int, value int) {
	node, e := this.knMap[key]
	if e { // 已经存在，改 + 升舱, 总个数不变
		node.value = value
		// 从旧得上面取下来
		oldLink := this.freqLinkMap[node.freq]
		oldLink.Pick(node)
		if node.freq == this.minFreq && oldLink.size == 0 { // 旧链已经空了
			this.minFreq++
		}

		node.freq++
		// 加到新的链头
		this.GetListByFreq(node.freq).InsertHead(node)
		return
	}
	// 原先没有的，要判断是否需要淘汰
	newNode := &Node{value: value, key: key, freq: 1}

	if this.size < this.cap { // 有空间，不必淘汰
		oneLink := this.GetListByFreq(1)
		oneLink.InsertHead(newNode)
		this.size++
	} else {
		// 驱逐旧的
		drop := this.freqLinkMap[this.minFreq].tail.prev
		this.freqLinkMap[this.minFreq].Pick(drop)
		delete(this.knMap, drop.key)
		this.GetListByFreq(1).InsertHead(newNode)
	}
	this.minFreq = 1
	this.knMap[key] = newNode
	return
}
