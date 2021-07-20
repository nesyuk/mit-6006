package main

import (
	"fmt"
)


// finds a peak in a row
func findRowGlobalMax(arr []int, idx int) (int, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("row is empty")
	}

	if idx > 0 && arr[idx] < arr[idx-1] {
		newLeft := idx / 2 
		return findRowGlobalMax(arr, newLeft)
	} else if idx < len(arr)-1 && arr[idx] < arr[idx+1] {
		newRight := (len(arr) + idx) / 2
		return findRowGlobalMax(arr, newRight)
	} else {
		return idx, nil
	}
}

func findPeak2D(arr [][]int, rowIdx int) (int, int, error) {
	if len(arr) == 0 {
		return 0, 0, fmt.Errorf("2D array is empty")
	}

	midColIdx := len(arr[rowIdx]) / 2
	colMax, err := findRowGlobalMax(arr[rowIdx], midColIdx)
	if err != nil {
		return 0, 0, err
	}

	if rowIdx > 0 && arr[rowIdx][colMax] < arr[rowIdx-1][colMax] {
		nextTop := rowIdx / 2
		return findPeak2D(arr, nextTop)
	} else if rowIdx < len(arr)-1 && arr[rowIdx][colMax] < arr[rowIdx+1][colMax] {
		nextBottom := (len(arr) + rowIdx) / 2
		return findPeak2D(arr, nextBottom)
	} else {
		return rowIdx, colMax, nil
	}
}

func FindPeak2D(arr [][]int) {
	fmt.Println("go!")
        midRowIdx := len(arr)/2
	row, col, err := findPeak2D(arr, midRowIdx)
	if err != nil {
		fmt.Printf("error: %v", err)
	} else {
		fmt.Printf("array: %v, col: %d, row: %d, peak: %v\n", arr, col, row, arr[row][col])
	}
}

func main() {
	FindPeak2D([][]int{{1, 2, 3}, {1, 2, 3}})
	FindPeak2D([][]int{{1, 2, 3, 4}, {14, 13, 12, 5}, {15, 9, 11, 17}, {16, 17, 19, 20}})
}
