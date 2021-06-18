package main

import (
	"fmt"
	"math/rand"
)

func main() {
	newRange := 401000000 - 56000000
	number := rand.Intn(newRange) + 56000000
	fmt.Println(number)
}