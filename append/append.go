package main

import (
	"fmt"
)

// a=[1 10 3], b=[1 10]
func main() {
	a := []int{1, 2, 3}
	b := append(a[:1], 10)
	fmt.Printf("a=%v, b=%v\n", a, b)
}
