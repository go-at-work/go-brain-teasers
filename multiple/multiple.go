package main

import (
	"fmt"
)

func main() {
	a, b := 1, 2
	b, c := 3, 4
	fmt.Println(a, b, c)
}

// Output: 1 3 4
// The code demonstrates that the second assignment to b shadows the first one, so the value of b in the output is 3, not 2.
// The variable c is declared in the same scope as b, so it can be used after the assignment.
