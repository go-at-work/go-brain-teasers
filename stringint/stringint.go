package main

import (
	"fmt"
)

func main() {
	i := 169
	s := string(i)
	fmt.Println(s)
}

// This code will print: ©
// It’ll treat this integer as arune. The rune 169 is the copyright sign (©).
