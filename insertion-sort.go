package main

import (
	"fmt"
)

// [5, 9, 8]

func sort(arr []int) {
	for start := 1; start < len(arr); start++ {
		for current := start; current > 0 && arr[current - 1] > arr[current]; current-- {
			arr[current], arr[current-1] = arr[current-1], arr[current];
		}
	}
}

func run() {
	arr := []int{0}
	fmt.Printf("from: %v\n", arr);
        sort(arr)
	fmt.Printf("to: %v\n", arr);

	arr = []int{0, 1}
	fmt.Printf("from: %v\n", arr);
        sort(arr)
	fmt.Printf("to: %v\n", arr);

	arr = []int{1, 0}
	fmt.Printf("from: %v\n", arr);
        sort(arr)
	fmt.Printf("to: %v\n", arr);

	arr = []int{1, 2, 3}
	fmt.Printf("from: %v\n", arr);
        sort(arr)
	fmt.Printf("to: %v\n", arr);

	arr = []int{3, 2, 1}
	fmt.Printf("from: %v\n", arr);
        sort(arr)
	fmt.Printf("to: %v\n", arr);
}


func main() {
	run()
}
