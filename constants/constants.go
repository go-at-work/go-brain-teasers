package main

import (
	"fmt"
)

// 3.1415929203539825
// π is a valid identifier,
// and that 355 / 113.0 actually compiles, the right side of the = are constants, not variables

func main() {
	var π = 355 / 113.0
	fmt.Println(π)
}
