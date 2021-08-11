package main

import (
	"fmt"
	"math/rand"
)

func insertionSort(a []float32) {
	for i := 0; i < len(a); i++ {
		for j := i; j > 0 && a[j-1] > a[j]; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}

func sort(a []float32) []float32 {
	buckets := make([][]float32, len(a))
	for _, num := range a {
		// determine a bucket number
		idx := int(num*float32(len(a)))
		// put in the bucket	
		buckets[idx] = append(buckets[idx], num)
	}
	i := 0
	for _, nums := range buckets {
		switch len(nums) {
			case 0:
				continue
			case 1:
				a[i] = nums[0]
			default:
				insertionSort(buckets[i])
				for j, num := range buckets[i] {
					a[i+j] = num
				}
		}	
		i += len(nums)
	}
	return a
}

func run() {
	fmt.Println("go!")
	rand.Seed(42)
	a := make([]float32, 0)
	for i := 10; i > 0; i-- {
		a = append(a, rand.Float32())
	}
	fmt.Println(a)
	a = sort(a)
	fmt.Println(a)
}

func main() {
	run()
}
