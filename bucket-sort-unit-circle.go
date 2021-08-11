package main

import (
	"fmt"
	"math"
	"time"
	"math/rand"
)

func sort(distances []float64) {
	buckets := make([][]float64, len(distances)*n)
	n := float64(len(distances))
	for _, d := range distances {
		//idx := int(d * n) // bad
		idx := int(d*d*n)
		buckets[idx] = append(buckets[idx], d)
	}
	sizes := make(map[int]int)
	for _, b := range buckets {
		sizes[len(b)]++
	}
	fmt.Println("sizes", sizes)
}

func run() {
	fmt.Println("go!")
	rand.Seed(time.Now().UnixNano())
	points := 10000
	distances := make([]float64, 0)
	for i := 0; i < points; i++ {
		x := rand.Float32()
		y := rand.Float32()
		r := math.Sqrt(float64(x*x) + float64(y*y))
		distances = append(distances, r)
	}
	sort(distances) 
}

func main() {
	run()
}
