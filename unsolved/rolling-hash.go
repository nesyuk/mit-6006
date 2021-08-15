package main

import (
	"fmt"
)

type RollingHash struct {
	hash int
	base int
	prime int
	capacity int // base ** size
}

func (h *RollingHash) Slide(old int, new int) {
	// hash = (n * base - old * base**size + new) mod prime // n is a string from which hash is calculated
	// hash = ((n*base mod prime) - (old * base**size mod prime) + new mod prime) mod prime
	// hash = (hash - (old * capacity mod prime) + new mod prime) mod prime
	// 
}

func (h *RollingHash) Append(s int) {
	h.hash = h.Hash(h.hash*h.base + h.Hash(s))
	h.capacity = h.Hash(h.capacity * h.base)
}

func (h *RollingHash) Skip(s int) {
	fmt.Println("capacity", h.capacity)
	h.capacity = h.Hash(h.capacity / h.base)
	//h.hash = h.Hash(h.hash - s*h.capacity + h.prime*h.hash) // add h.p * h.base to avoid negative sum h.hash - s*h.capacity
	h.hash = h.Hash(h.hash - s*(h.capacity/h.base)) // add h.p * h.base to avoid negative sum h.hash - s*h.capacit
}

func (h *RollingHash) Hash(s int) int {
	return s % h.prime
}

func NewRollingHash() RollingHash {
        // base: 143859: number of symbols in Unicode (number of runes), prime: 2147483869
	return RollingHash{hash: 0, base: 100, prime: 23, capacity: 1}
}

func run() {
	fmt.Println("go!")
	rh := NewRollingHash()
        //str := []rune{"abracatabra"}
	//search := "cat"
	rh.Append(3)
	fmt.Println(rh.hash)
	rh.Append(14)
	fmt.Println(rh.hash)
	rh.Append(15)
	fmt.Println(rh.hash)
	rh.Append(92)
	fmt.Println(rh.hash)
	rh.Append(65)
	fmt.Println(rh.hash)
	rh.Skip(3)
	fmt.Println(rh.hash)
	rh.Append(35)
	fmt.Println(rh.hash)
}

func main() {
	run()
}
