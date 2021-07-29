package main

import (
	"fmt"
	"math/rand"
	"time"
)

func shuffle(a *[]int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(*a); i++ {
		r := rand.Intn(len(*a))
		(*a)[i], (*a)[r] = (*a)[r], (*a)[i]
	}
}

func run() {
	fmt.Println("go!")

	arr := []int{1, 2, 3}
	fmt.Printf("before: %v\n", arr)
	shuffle(&arr)
	fmt.Printf("after: %v\n", arr)

	arr = []int{1, 2, 3}
	shuffle(&arr)
	fmt.Printf("after: %v\n", arr)
}

func main() {
	run()
}
