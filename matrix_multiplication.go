package main

import (
	"fmt"
)


// 1 2   1 2 3  
// 3 4   4 5 6
// 5 6

// 1 * 1 + 1 * 2 + 1 * 3

func mutiply(x [][]int, y [][]int) (*[][]int, error) {
	// x: m x n
	m := len(x)
	n := len(x[0])

	
        // y: n x l
        if len(y) != n {
		return nil, fmt.Errorf("matrices are incompatible")
	}
	l := len(y[0])

	// z: m x l
	z := make([][]int, m)
	for i := range z {
		z[i] = make([]int, l)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < l; j++ {
			for k := 0; k < m; k++ {
				z[i][j] += x[i][k] * y[k][j]
			} 
		}
	}
	return &z, nil
}

func print(m [][]int) {
	for i := 0; i < len(m); i++ {
		fmt.Println(m[i]);
	}
}

func run() {
	fmt.Println("go!")
	x := [][]int{{1, 2}, {3, 4}}
	y := [][]int{{5, 6}, {7, 8}}
	z, err := mutiply(x, y)
	if err != nil {
		fmt.Println(err)
	} else {
		print(*z)
	}
}

func main() {
	run()
}
