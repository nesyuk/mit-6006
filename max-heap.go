package main

import (
	"fmt"
)

type MaxHeap struct {
	a []int
	length int
}

func (h *MaxHeap) maxHeapify(i int) {
	for root := i; root > 0; root /= 2 {
		left := root*2
		right := root*2+1
		if left <= h.length && h.a[left] > h.a[root] {
			h.a[root], h.a[left] = h.a[left], h.a[root]
		}
		if right <= h.length && h.a[right] > h.a[root] {
			h.a[root], h.a[right] = h.a[right], h.a[root]
		}
	}
}

func (h MaxHeap) Sort(a []int) []int {
	hp := h.build(a)
	for hp.GetMax() != nil {
		hp.ExtractMax()	
	}
	return hp.a[1:]
}

func (h MaxHeap) build(a []int) *MaxHeap {
	h.length = len(a)
	h.a = make([]int, len(a)+1)
	copy(h.a[1:], a[:])
	for idx := h.length / 2; idx > 0; idx-- {
		h.maxHeapify(idx)
	}
	return &h
}

func (h MaxHeap) GetMax() *int {
	if h.length == 0 {
		return nil
	}
	m := h.a[1]
	return &m
}

func (h *MaxHeap) ExtractMax() *int {
	if h.length == 0 {
		return nil
	}
	last := h.length
	h.a[1], h.a[last] = h.a[last], h.a[1]
	h.length = last - 1
	h.maxHeapify(h.length/2)
	m := h.a[last]
	return &m
}

func run() {
	fmt.Println("go!")
	var heap MaxHeap
	hp := heap.build([]int{1,2,3,4,5})
	fmt.Printf("max: %v\n", *hp.ExtractMax())
	fmt.Printf("max: %v\n", *hp.ExtractMax())
	fmt.Printf("max: %v\n", *hp.ExtractMax())
	fmt.Printf("max: %v\n", *hp.ExtractMax())
	fmt.Printf("max: %v\n", *hp.ExtractMax())

	hp = heap.build([]int{5, 4, 3, 2, 1})
	fmt.Println("-------")
	fmt.Printf("max: %v\n", *hp.ExtractMax())
	fmt.Printf("max: %v\n", *hp.ExtractMax())
	fmt.Printf("max: %v\n", *hp.ExtractMax())
	fmt.Printf("max: %v\n", *hp.ExtractMax())
	fmt.Printf("max: %v\n", *hp.ExtractMax())

	hp = heap.build([]int{2, 3, 4, 6, 3, 4, -1, 0})
	fmt.Println("-------")
	fmt.Printf("max: %v\n", *hp.ExtractMax())
	fmt.Printf("max: %v\n", *hp.ExtractMax())
	fmt.Printf("max: %v\n", *hp.ExtractMax())
	fmt.Printf("max: %v\n", *hp.ExtractMax())
	fmt.Printf("max: %v\n", *hp.ExtractMax())
	fmt.Printf("max: %v\n", *hp.ExtractMax())
	fmt.Printf("max: %v\n", *hp.ExtractMax())
	fmt.Printf("max: %v\n", *hp.ExtractMax())

	fmt.Println("-------")
	fmt.Printf("sorted: %v\n", heap.Sort([]int{2, 3, 4, 6, 3, 4, -1, 0}))
}

func main() {
	run()
}
