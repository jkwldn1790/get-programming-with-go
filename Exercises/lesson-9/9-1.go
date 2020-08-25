package main

import (
	"fmt"
)

// You can use len(variable) to measure the length of a string
// 'A' and "A" are not the same thing ... 'A' = rune and "A" is string
// Byte is an alias for the uint8 type and rune is an alias for int32

func main() {
	message := "HELLO WELCOME"
	var cipher byte = 1
	fmt.Printf("Length of message: %v\n", len(message))
	for i := 0; i < len(message); i++ {
		c := message[i] + cipher
		fmt.Printf("%c\n", c)
	}
}
