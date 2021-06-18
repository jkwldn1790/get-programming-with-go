package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	distanceToMars = 62100000
	secondsPerDay = 86400
)

func main() {
	fmt.Printf("%-20v %-10v %-20v %s\n", "Spaceline", "Days", "Trip-Type", "Price")
	header := []string{strings.Repeat("=", 58)}
	fmt.Println(strings.Join(header, ""))
	for count := 10; count > 0; count-- {

		var spaceline string

		speed := rand.Intn(15) + 16

		days := distanceToMars / speed / secondsPerDay

		ticketPrice := 20 + speed

		switch rand.Intn(3) {
		case 0:
			spaceline = "SpaceX"
		case 1:
			spaceline = "Virgin Galactic"
		case 2:
			spaceline = "Blue Origin"
		}

		fmt.Printf("%-20v %-10v %-20v $ %v\n", spaceline, days, "One-Way", ticketPrice)
		fmt.Printf("%-20v %-10v %-20v $ %v\n", spaceline, (days * 2), "Round-Trip", (ticketPrice * 2))
		
	}
}