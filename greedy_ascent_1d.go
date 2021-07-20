package main

import (
	"fmt"
	"log"
)

// definition: Xi is a peak if Xi >= Xi-1 and Xi >= Xi+1

func findPeak(arr []int, idx int) (int, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("array is empty")
	}
	if idx > 0 && arr[idx] < arr[idx-1] {
		midLeft := idx/2
		return findPeak(arr, midLeft)
	} else if idx < len(arr) - 1 && arr[idx] < arr[idx+1] {
		midRight := (len(arr) + idx) / 2
		return findPeak(arr, midRight)
	} else {
		return idx, nil
	}
}

func run(arr []int) {
        idx := len(arr)/2
	idx, err := findPeak(arr, idx) 
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("peak for %v is %d for index: %d\n", arr, arr[idx], idx)
}

func main() {
	fmt.Println("go!")
        arr := []int{1}
	run(arr)

	arr = []int{1, 2}
	run(arr)

	arr = []int{2, 1}
	run(arr)

	arr = []int{1, 2, 1}
	run(arr)

	arr = []int{4, 2, 3}
	run(arr)

	arr = []int{3, 2, 4}
	run(arr)

	arr = []int{1, 2, 3, 4}
	run(arr)
}
