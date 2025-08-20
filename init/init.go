package main

import (
	"fmt"
)

func init() {
	fmt.Printf("A ")
}
func init() {
	fmt.Print("B ")
}
func main() {
	fmt.Println()
}

// Output: A B
// The init functions are called in the order they are defined in the file.
// One of the problems with package-level variables or init is that you can’t return
// error values. If you have an error in init, the best course of action is to panic
// and not continue execution in a bad state. To support that, some of the Go
// packages provide a Must version of their New functions
