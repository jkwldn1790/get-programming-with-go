package main

import "fmt"

func main() {
	const lightSpeed = 299792 // km/s
	const hoursPerDay = 24
	const turtleSpeed = 100800 // km/h
	var distance = 56000000 // km

	fmt.Println(distance/lightSpeed, "seconds")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")

	distance = 96300000
	fmt.Println(distance/turtleSpeed/hoursPerDay, "days")
}