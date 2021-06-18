package main

import "fmt"

func main() {
	fmt.Printf("%-15v $%5v\n", "SpaceX", 94) // -15 is saying to reserve 15 spaces if you can. If string is longer than the fifteen characters then it will proceed.
	fmt.Printf("%-15v $%5v\n", "Virgin Galactic", 100) // %4v is saying that there MUST be 4 spaces of padding in between.
}