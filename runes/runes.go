package main

import (
	"fmt"
)

func main() {
	msg := "π = 3.14159265358..."
	fmt.Printf("%T ", msg[0])
	for _, c := range msg {
		fmt.Printf("%T\n", c)
		break
	}
}

// uint8 int32
// Go strings are UTF-8 encoded. When you access a string with [] or with len,
// Go will access the underlying []byte. Slice is a type in Go that’s aliased to uint8
