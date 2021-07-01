package main

import (
	"fmt"
)


type fahrenheit float64
type celsius float64
type kelvin float64

// convert farenheit to celsius

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32) * 5/9)
}

// convert kelvin to celsius

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

// convert farenheit to kelvin

func (f fahrenheit) kelvin() kelvin {
	return kelvin((f - 32) * 5/9 + 273.15)
}

// convert celsius to fahrenheit

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9/5) + 32)
}

// convert celsius to kelvin

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

// convert kelvin to fahrenheit

func (k kelvin) fahrenheit() fahrenheit {
	return fahrenheit((k - 273.15) * 9/5 + 32)
}

func main() {
	var temp fahrenheit = 0
	fmt.Printf("%v\n", kelvin.celsius(kelvin(temp)))
	fmt.Printf("%.2f\n", fahrenheit.celsius(temp)) // -17.778
	fmt.Printf("%.4f\n", fahrenheit.kelvin(temp)) // 255.372
	fmt.Printf("%.4f\n", celsius.fahrenheit(celsius(temp))) // 32.0000
	fmt.Printf("%.4f\n", celsius.kelvin(celsius(temp))) // 273.15
	fmt.Printf("%.4f\n", kelvin.fahrenheit(kelvin(temp))) // 273.15
}