package main

import (
	"fmt"
	b "math/big"
)

func run() {
	prime := b.NewInt(23)
	num := b.NewInt(6)
	inv := b.NewInt(1)
	fmt.Println(inv.ModInverse(num, prime))

	// base: 143859, prime: 2147483869
	prime = b.NewInt(2147483869)
	base := b.NewInt(143859)
	fmt.Println(inv.ModInverse(base, prime))

	fmt.Println(base)
	base = base.Mul(base, base)
	base = base.Mod(base, prime)
	fmt.Println(base)
	base.ModInverse(base, prime)
	fmt.Println(base)
}

func main() {
	run()
}
