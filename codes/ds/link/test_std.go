package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	e1 := l.PushBack(1)
	e2 := l.PushFront(2)
	l.InsertAfter(3, e1)
	l.Remove(e2)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
