package main

import (
	"fmt"
)

// 1.2100000000000002
// in floating points, we sacrifice accuracy for speed
// If you run the same code in Python, Java, C … you will see the same output
// check for “roughly equal” and decide what is an acceptable threshold.
// Testing libraries such as stretchr/testify have ready-made functions (such as InDelta) to check that two floating numbers are approximately equal.
// If you need better accuracy, look into math/big or external packages such as shopspring/decimal.
func main() {
	n := 1.1
	fmt.Println(n * n)
}
