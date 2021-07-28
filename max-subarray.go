package main

import (
	"fmt"
)

const MIN_INT = -1000000 //


func findMax(arr []int, l int, r int) (int, int, int) {
	if l == r {
		return l, r, arr[l]
	}
	mid := (l + r) / 2
	ll, lr, lsum := findMax(arr, l, mid)
	rl, rr, rsum := findMax(arr, mid+1, r)
	cl, cr, csum := findMaxCrossing(arr, l, mid, r)
	
	if lsum >= rsum && lsum >= csum {
		return ll, lr, lsum
	} else if rsum >= lsum && rsum >= csum {
		return rl, rr, rsum
	} else {
		return cl, cr, csum
	}
}

func findMaxCrossing(arr []int, l int, mid int, r int) (int, int, int) {
	lsum := MIN_INT
	lp := mid
	sum := 0
	for i := lp; i >= l; i-- {
		sum += arr[i]
		if sum > lsum {
			lp = i
			lsum = sum
		}
	}

	rsum := MIN_INT
	rp := mid + 1
	sum = 0
	for i := rp; i <= r; i++ {
		sum += arr[i]
		if sum > rsum {
			rp = i
			rsum = sum
		}		
	}
	return lp, rp, lsum + rsum
}

func findMaxSubarray(arr []int) (int, int, int) {
	var diffs []int 
	for i := 1; i < len(arr); i++ {
		diffs = append(diffs, arr[i] - arr[i-1]) 
	}
	l, r, sum := findMax(diffs, 0, len(arr)-2) // diffs has len - 1
	return l+1, r+1, sum
}

func run() {
	l, r, sum := findMaxSubarray([]int{100, 113, 110, 85, 105, 102, 86, 63, 81, 101, 94, 106, 101, 79, 94, 90, 97})
	fmt.Printf("from: %d, to: %d, sum: %d\n", l, r, sum)
	
	l, r, sum = findMaxSubarray([]int{100, 90, 80, 70, 60, 50, 40, 30, 20, 10, 0}) // strictly declining
	fmt.Printf("from: %d, to: %d, sum: %d\n", l, r, sum)

	l, r, sum = findMaxSubarray([]int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}) // strictly rising
	fmt.Printf("from: %d, to: %d, sum: %d\n", l, r, sum)
}

func main() {
	run()
}
