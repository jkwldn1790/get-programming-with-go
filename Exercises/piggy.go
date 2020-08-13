package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sum := 20.0
	Nickels, Dimes, Quarters := 0, 0, 0
	for sum > 0.0 {
		curr := rand.Intn(3)
		if curr == 0 {
			sum -= 0.05
			Nickels += 1
		} else if curr == 1 {
			sum -= 0.10
			Dimes += 1
		} else {
			sum -= 0.25
			Quarters += 1
		}
	}
	fmt.Printf("\nNickels: %v\nDimes: %v\nQuarters: %v\n", Nickels, Dimes, Quarters)
}