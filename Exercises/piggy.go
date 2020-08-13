package main

import (
	"fmt"
	"math/rand"
)

// func chg() {
// 	curr := rand.Intn(3)
// 	coin := 0.0
// 	switch curr {
// 	case 0:
// 		coin := 0.05
// 	case 1:
// 		coin := 0.10
// 	case 2:
// 		coin := 0.25
// 	}
// 	return coin
// }

func main() {
	curr := rand.Intn(3)
	coin := 0.0
	sum := 0
	for i := 0; sum < 20.0; i += sum {
		if curr == 0 {
			coin = 0.05
		} else if curr == 1 {
			coin = 0.10
		} else {
			coin = 0.25
		}
	}
	fmt.Printf("%v test", coin)
}