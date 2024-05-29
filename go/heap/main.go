package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type MaxHeap struct {
	sort.IntSlice
}

func (mh *MaxHeap) Less(i, j int) bool { return mh.IntSlice[i] > mh.IntSlice[j] }
func (mh *MaxHeap) Push(x any) {
	mh.IntSlice = append(mh.IntSlice, x.(int))
}
func (mh *MaxHeap) Pop() any {
	n := len(mh.IntSlice)
	x := mh.IntSlice[n-1]
	mh.IntSlice = mh.IntSlice[:n-1]
	return x
}

func main() {
	mh := &MaxHeap{sort.IntSlice{2, 1, 5}}
	heap.Init(mh)
	heap.Push(mh, 8)
	fmt.Printf("max: %v\n", mh.IntSlice[0])

	heap.Pop(mh)
	fmt.Printf("max: %v\n", mh.IntSlice[0])
}
