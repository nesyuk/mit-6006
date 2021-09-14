package main

import "fmt"

/// 6, 7, 0, 1, 2, 3, 4, 5


func find(nums []int, target int) int {
	left := 0
	right := len(nums)-1

	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] {
			// left part is sorted

			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// right part is sorted

			if target <= nums[right] && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	idx := find([]int{3, 4, 5, 6, 0, 1, 2}, 6)
	fmt.Println("expected: %v, got: %v", 3, idx)

	idx = find([]int{3, 4, 5, 6, 0, 1, 2}, 3)
	fmt.Println("expected: %v, got: %v", 0, idx)

	idx = find([]int{3, 4, 5, 6, 0, 1, 2}, 0)
	fmt.Println("expected: %v, got: %v", 4, idx)
}
