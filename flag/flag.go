package main

import (
	"fmt"
)

type Flag int

func main() {
	var i interface{} = 3
	f := i.(Flag)
	fmt.Println(f)
}

// This code will panic.
// The type assertion i.(Flag) will fail because the underlying type of i is int, not Flag.
