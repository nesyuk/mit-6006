package main

import (
	"fmt"
)

type MaxQueue struct {
	a []int 
	length int
}

func Build(a []int) *MaxQueue {
	queue := MaxQueue{a: make([]int, len(a) + 1)}
	for _, el := range a {
		queue.Insert(el)
	}
	return &queue
}

func (q *MaxQueue) GetMax() int {
	return q.a[1]
}

func (q *MaxQueue) Insert(el int) {
	q.length = q.length + 1
	q.a[q.length] = el
	for i := q.length; i >= 2 && q.a[i/2] < q.a[i]; i/=2 {
		q.a[i/2], q.a[i] = q.a[i], q.a[i/2]
	}
}

func (q *MaxQueue) IncreaseKey(el, key int) error {
	if key < q.a[el] {
		return fmt.Errorf("current key is higher than the new key")
	}
	q.a[el] = key
	for i := el; i >= 2 && q.a[i/2] < q.a[i]; i/=2 {
                q.a[i/2], q.a[i] = q.a[i], q.a[i/2]
        }
	return nil
}

func (q *MaxQueue) maxHeapify(parent int) {
	left := parent*2
	largest := parent
	if left <= q.length && q.a[left] > q.a[parent] {
		largest = left
	}
	right := left + 1
	if right <= q.length && q.a[right] > q.a[largest] {
		largest = right
	}

	if largest != parent {
		q.a[parent], q.a[largest] = q.a[largest], q.a[parent]
		q.maxHeapify(largest)
	}
}

func (q *MaxQueue) ExtractMax() int {
	m := q.a[1]

	q.a[1], q.a[q.length] = q.a[q.length], q.a[1]
	q.length = q.length - 1
	q.a = q.a[:len(q.a) - 1] // delete last

	q.maxHeapify(1)
	return m
}

func run() {
	fmt.Println("go!")

	pq := Build([]int{1, 2, 3})
	fmt.Printf("%d\n", pq.ExtractMax())
	fmt.Printf("%d\n", pq.ExtractMax())
	fmt.Printf("%d\n", pq.ExtractMax())
	
	pq = Build([]int{3, 2, 1})
	fmt.Printf("%d\n", pq.ExtractMax())
	fmt.Printf("%d\n", pq.ExtractMax())
	fmt.Printf("%d\n", pq.ExtractMax())

	pq = Build([]int{2, 3, 1})
	fmt.Printf("%d\n", pq.ExtractMax())
	fmt.Printf("%d\n", pq.ExtractMax())
	fmt.Printf("%d\n", pq.ExtractMax())
}

func main() {
	run()
}
