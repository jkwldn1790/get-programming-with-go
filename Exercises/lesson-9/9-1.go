package main

import (
	"fmt"
)

// How do you measure the length of a string

func main() {
	message := "HELLO WELCOME"
	var cipher byte = 1
	fmt.Printf("Length of message: %v\n", len(message))
	for i := 0; i < len(message); i++ {
		c := message[i] + cipher
		fmt.Printf("%c\n", c)
	}
}
