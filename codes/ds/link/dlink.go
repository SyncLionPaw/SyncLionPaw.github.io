package main

import "errors"

type Node struct {
	// value ...
	prev  *Node
	next  *Node
	value int
}

type DoubleLinkList struct {
	head *Node
	tail *Node
	size int
}

func NewDoubleLinkList() *DoubleLinkList {
	h, t := &Node{nil, nil, 0}, &Node{nil, nil, 0}
	h.next = t
	t.prev = h
	return &DoubleLinkList{h, t, 0}
}

// pick up node from list
func (d *DoubleLinkList) Pick(node *Node) error {
	// assert node in list and node isn't head and tail
	if node == nil {
		return nil
	}
	if node == d.head || node == d.tail {
		return errors.New("sentinel node was protected, should not be pick.")
	}
	if p, n := node.prev, node.next; p == nil || n == nil {
		return errors.New("node you've offerred was not in list.")
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	d.size--
	return nil
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

func main() {
	dll := NewDoubleLinkList()
	// 创建节点
	n1 := &Node{value: 1}
	n2 := &Node{value: 2}
	n3 := &Node{value: 3}

	// 插入节点
	dll.InsertHead(n1)
	dll.InsertTail(n2)
	dll.InsertHead(n3)

	// 打印链表内容
	cur := dll.head.next
	for cur != dll.tail {
		println(cur.value)
		cur = cur.next
	}

	// 删除 n1
	err := dll.Pick(n1)
	if err != nil {
		println("Pick error:", err.Error())
	}

	// 再次打印链表内容
	cur = dll.head.next
	for cur != dll.tail {
		println(cur.value)
		cur = cur.next
	}
}
