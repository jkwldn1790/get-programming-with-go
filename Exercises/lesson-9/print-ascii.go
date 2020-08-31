package main

import (
	"fmt"

	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	for i := 0; i < 250; i++ {
		fmt.Printf("Integer %v string-literal %[1]c\n", i)
	}
}
