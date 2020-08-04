package main

import (
	"fmt"
	"math/rand"
)

var (
	departDate = "October 13, 2020"
	destination = "Mars"
	distance = 62100000
	tickets = 10
)

const (
	secDay = 86400
	priceExp = 2
)

func main() {
	fmt.Printf("\nDeparting on: %v\n\nDestination: %v\nDistance: %v km\n\n", departDate, destination, distance)
	fmt.Println("Spaceline         Speed  Days  TripType      Price")
	fmt.Println("==================================================")
	for tickets > 0 {		
		tripType := "One-way"
		spaceLineSelect := "SpaceX"
		spaceLine := rand.Intn(3)
		speed := rand.Intn(14) + 16
		days := (distance / speed) / secDay
		price := speed * priceExp
		switch spaceLine {
		case 0:
			spaceLineSelect = "Space Adventures"
		case 1:
			spaceLineSelect = "SpaceX"
		case 2:
			spaceLineSelect = "Virgin Galactic"
		}
		if days > 30 {
			tripType = "Round-trip"
			price = price * 2
		}
		fmt.Printf("%-18v %3v %5v %-16v $%3v\n", spaceLineSelect, speed, days, tripType, price)
		tickets--
	}
}