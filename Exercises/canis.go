package main

import (
	"fmt"
)

const canisDistance = 236000000000000000
const lightSpeed = 299792 // m/s
const secPerYear = 31536000
const secondsPerDay = 86400
const daysPerYear = 365

const years = canisDistance / lightSpeed / secondsPerDay / daysPerYear

func main() {
	fmt.Printf("Light Years to Canis: %v\n", years)
}