package main

import (
	"fmt"
	"math/big"
)

func main() {
	/* Setting these using NewInt because when using
	the "math/big" package all integers need to be in
	the big package.
	*/
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)

	distance := new(big.Int)
	// Reads in the integer as a string using base 10
	distance.SetString("2000000000000000000", 10)
	fmt.Println("Alpha Centauri is", distance, "km away.")
	
	seconds := new(big.Int)
	seconds.Div(seconds, secondsPerDay)

	days := new(big.Int)
	days.Div(distance, lightSpeed)

	fmt.Println("That is", days, "days of travel at light speed.")
}