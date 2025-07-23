package main

import (
	"container/heap"
	"fmt"
)

type Info struct {
	s   string
	val int
}

type InfoHeap []Info

func (ih InfoHeap) Len() int {
	return len(ih)
}

func (ih InfoHeap) Less(i, j int) bool {
	return ih[i].val < ih[j].val
}

func (ih InfoHeap) Swap(i, j int) {
	ih[i], ih[j] = ih[j], ih[i]
}

func (ih *InfoHeap) Push(x any) {
	*ih = append(*ih, x.(Info))
}

func (ih *InfoHeap) Pop() any {
	n := len(*ih)
	x := (*ih)[n-1]
	(*ih) = (*ih)[:n-1]
	return x
}

func main() {
	ih := &InfoHeap{Info{"a", 2}, Info{"b", 3}}
	heap.Init(ih)
	ans := heap.Pop(ih)
	fmt.Printf("ans: %v\n", ans)
}