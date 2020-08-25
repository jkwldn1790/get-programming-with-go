package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 250; i++ {
		fmt.Printf("Integer %v string-literal %[1]c\n", i)
	}
}
