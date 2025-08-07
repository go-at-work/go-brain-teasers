package main

import (
	"fmt"
)

// 7
// Go strings are UTF-8 encoded, the rune ó is taking 2 bytes; hence, the total length of the string is 7.
func main() {
	city := "Kraków"
	fmt.Println(len(city))
}
