package main

import (
	"fmt"
)

// This code will print: 0
// the key "errors" is not found in the map and you’ll get back the zero value for int, which is 0.
func main() {
	var m map[string]int
	fmt.Println(m["errors"])
}
