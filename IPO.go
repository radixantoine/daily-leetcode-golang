package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Pair struct {
	profit  int
	capital int
}

type projectHeap []int

func (h projectHeap) Len() int           { return len(h) }
func (h projectHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h projectHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *projectHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *projectHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	p := make([]Pair, n)
	for i := 0; i < n; i++ {
		p[i] = Pair{profits[i], capital[i]}
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].capital < p[j].capital
	})

	j := 0

	maxHeap := &projectHeap{}
	heap.Init(maxHeap)

	for ; k > 0; k-- {
		for j < n && p[j].capital <= w {
			heap.Push(maxHeap, p[j].profit)
			j++
		}
		if maxHeap.Len() == 0 {
			return w
		}
		w += heap.Pop(maxHeap).(int)
	}
	return w
}

func main() {
	fmt.Println(
		findMaximizedCapital(2, 0, []int{1, 2, 3}, []int{0, 1, 1}))
}
