package graph

import (
	"container/heap"
)

type Item struct {
	Node NodeID
	Dist float64
	Index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	// returns the length
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// returns true if item i has higher priority than j
	return pq[i].Dist < pq[j].Dist
}

func (pq PriorityQueue) Swap(i, j int) {
	// swap elements and update their index fields
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = j
	pq[j].Index = i
}

func (pq *PriorityQueue) Push(x any) {
	// cast x to *Item, append to slice, set Index
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	// remove last element fix indices return the item
	old := *pq
	n := len(*pq)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0: n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, dist float64) {
	// update item.Dist and call heap.Fix(pq, item.Index)
	item.Dist = dist
	heap.Fix(pq, item.Index)
}