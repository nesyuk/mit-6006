package main

import (
	"fmt"
)

func sort(a []int, k int) []int {
	counts := make([]int, k+1)
	
	// calculate counts
	// for ex, for [3, 4, 2, 1, 1, 3, 6] : [0, 2, 1, 2, 1, 0, 1]
	for _, e := range a {
		counts[e]++
	}
	
	// calculate number of elements less than or equal to i
	// for ex, for [0, 2, 1, 2, 1, 0, 1]: [0, 2, 3, 5, 6, 6, 7]
	for i := 1; i < k; i++ {
		counts[i] = counts[i] + counts[i-1]
	}

	sorted := make([]int, len(a))
	for _, e := range a {
		sorted[counts[e-1]] = e
		counts[e-1]++
	}
	return sorted
}

func run() {
	a := []int{3, 4, 2, 1, 1, 3, 6}
	k := 6
	sorted := sort(a, k)
	fmt.Println(a, sorted)
}

func main() {
	run()
}
