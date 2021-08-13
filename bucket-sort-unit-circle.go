package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

func distance(p []float64) float64 {
	return p[0]*p[0] + p[1]*p[1]
}

func insertionSort(points [][]float64) {
	for i := 1; i < len(points); i++ {
		for j := i; j > 0 && distance(points[i-1]) > distance(points[i]); j-- {
			points[i-1], points[i] = points[i], points[i-1]
		}
	}
}

func sort(points [][]float64) {
	buckets := make([][][]float64, len(points))
	n := float64(len(points))
	for _, p := range points {
		r := p[0]*p[0] + p[1]*p[1] // radius squared
		idx := int(r*n) // floor
		buckets[idx] = append(buckets[idx], p)
	}
        /*  test bucket sizes
		sizes := make(map[int]int)
		for _, b := range buckets {
			sizes[len(b)]++
		}
		fmt.Println("sizes", sizes)
	*/
	// sort each bucket with insertion sort
	for _, bucket := range buckets {
		insertionSort(bucket)		
	}
	i := 0
	for _, bucket := range buckets {
		for _, p := range bucket {
			points[i] = p
			i++
		}	
	}
}

func run() {
	fmt.Println("go!")
	rand.Seed(time.Now().UnixNano())
	n := 10
	points := make([][]float64, 0)
	for i := 0; i < n; {
		x := rand.Float64()
		y := rand.Float64()
		if math.Sqrt(x*x + y*y) <= 1.0 {
			points = append(points, []float64{x, y})
			i++
		}
	}
	sort(points) 
	for _, p := range points {
		fmt.Printf("x: %.5f, y: %.5f r: %.5f\n", p[0], p[1], math.Sqrt(p[0]*p[0] + p[1]*p[1]))
	}
}

func main() {
	run()
}
