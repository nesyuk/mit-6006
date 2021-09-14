package main

import "fmt"

func findPivot(nums []int) int {
	return binarySearch(nums, 0, len(nums)-1)
}

func binarySearch(nums []int, left, right int) int {
	for left < right {
		mid := left + (right - left) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		}
	}
	return left
}

func main() {
	pivot := findPivot([]int{3, 4, 5, 0, 1, 2})
	fmt.Println("expected: %v, got: %v", 3, pivot)

	pivot = findPivot([]int{3, 0, 1, 2})
	fmt.Println("expected: %v, got: %v", 1, pivot)

	pivot = findPivot([]int{3, 4, 5, 6, 2})
	fmt.Println("expected: %v, got: %v", 4, pivot)

	pivot = findPivot([]int{2, 3, 4, 5, 6})
        fmt.Println("expected: %v, got: %v", 0, pivot)
}
