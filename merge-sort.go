package main

import (
	"fmt"
)

func sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := sort(arr[:mid])
	right := sort(arr[mid:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	merged := make([]int, 0)

	l := 0
	r := 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			merged = append(merged, left[l]); l++
		} else {
			merged = append(merged, right[r]); r++
		}
	}

	for l < len(left) {
		merged = append(merged, left[l]); l++
	}

	for r < len(right) {
		merged = append(merged, right[r]); r++
	}
	return merged
}

func run() {
	arr := []int{0}
	fmt.Printf("from: %v\n", arr);
        sorted := sort(arr)
	fmt.Printf("to: %v\n", sorted);

	arr = []int{0, 1}
	fmt.Printf("from: %v\n", arr);
        sorted = sort(arr)
	fmt.Printf("to: %v\n", sorted);

	arr = []int{1, 0}
	fmt.Printf("from: %v\n", arr);
        sorted = sort(arr)
	fmt.Printf("to: %v\n", sorted);

	arr = []int{1, 2, 3}
	fmt.Printf("from: %v\n", arr);
        sorted = sort(arr)
	fmt.Printf("to: %v\n", sorted);

	arr = []int{3, 2, 1}
	fmt.Printf("from: %v\n", arr);
        sorted = sort(arr)
	fmt.Printf("to: %v\n", sorted);
}


func main() {
	run()
}
