/*Write a program that converts strings
to booleans*/
package main

import (
	"fmt"
)

var (
	eval bool
)

func main() {
	input := fmt.Sprintf("1")
	switch input {
	case "true", "yes", "1":
		eval = true
	case "false", "no", "0":
		eval = false
	default:
		fmt.Printf("Error:\n%v is not valid.", input)
	}
	fmt.Printf("%T\n%[1]v\n", eval)
}
