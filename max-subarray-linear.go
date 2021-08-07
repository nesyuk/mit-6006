package main

import (
	"fmt"
)


func findMax(arr []int) (low int, high int, maxSum int) {
	if len(arr) <= 1 {
		return low, high, maxSum
	}
	maxSum = arr[0]
	sum := maxSum
        currentLow, currentHigh := 0, 0

	for i := 1; i < len(arr); i++ {
		if sum > 0 {
			// keep adding
			sum += arr[i]
			currentHigh = i
		} else {
			// reset
			sum = arr[i]
			currentLow, currentHigh = i, i
		}

		if sum > maxSum {
			maxSum = sum
			low, high = currentLow, currentHigh
		}
	}
	return low, high, maxSum
}

func findMaxSubarray(arr []int) (int, int, int) {
	var diffs []int 
	for i := 1; i < len(arr); i++ {
		diffs = append(diffs, arr[i] - arr[i-1]) 
	}
	fmt.Println(diffs)
	l, r, sum := findMax(diffs) // diffs has len - 1
	return l+1, r+1, sum
}

func run() {
	l, r, sum := findMaxSubarray([]int{100, 113, 110, 85, 105, 102, 86, 63, 81, 101, 94, 106, 101, 79, 94, 90, 97})
	fmt.Printf("from: %d, to: %d, sum: %d\n", l, r, sum)
	
	l, r, sum = findMaxSubarray([]int{100, 90, 80, 70, 60, 50, 40, 30, 20, 10, 0}) // strictly declining
	fmt.Printf("from: %d, to: %d, sum: %d\n", l, r, sum) // 90-100 = -10

	l, r, sum = findMaxSubarray([]int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}) // strictly rising
	fmt.Printf("from: %d, to: %d, sum: %d\n", l, r, sum) // 100 - 0 = 100
}

func main() {
	run()
}
