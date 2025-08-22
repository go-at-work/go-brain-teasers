package main

import (
	"fmt"
	"time"
)

type Log struct {
	Message   string
	time.Time //  a field with no name, just a type. This is called embedding, it means that the Log type has all the methods and fields that time.Time has
	// This is useful for logging timestamps without needing to define a separate field for the time.
	// You can also embed interfaces in Go.
}

// 2009-11-10 00:00:00 +0000 UTC
// The %v verb will print all the struct fields.
func main() {
	ts := time.Date(2009, 11, 10, 0, 0, 0, 0, time.UTC)
	log := Log{"Hello", ts}
	fmt.Printf("%v\n", log)
}
