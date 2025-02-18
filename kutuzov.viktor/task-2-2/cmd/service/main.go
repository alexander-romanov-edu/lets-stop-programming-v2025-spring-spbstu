package main

import (
	"container/heap"
	"fmt"
)

// IntHeap представляет собой минимальную кучу для целых чисел.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return (*h)[0]
}

func main() {
	var N, k int
	fmt.Scan(&N)

	nums := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Scan(&k)

	result := findKthLargest(nums, k)
	fmt.Println(result)
}
