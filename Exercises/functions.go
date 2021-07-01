package main

import (
	"fmt"
)

type celsius float64
type kelvin float64
type farenheit float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func celsiusToFarenheit(c celsius) farenheit {
	return farenheit((c * 9.0/5.0) + 32.0)
}

func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func main() {
	var kel kelvin = 294.0
	cel := kelvinToCelsius(kel)
	far := celsiusToFarenheit(cel)

	var sunlitMoon celsius = 127.0

	sunlitMoonKelvin := celsiusToKelvin(sunlitMoon)

	fmt.Printf("%.2f °C %[1]T\n", cel)
	fmt.Printf("%.2f °F %[1]T\n", far)

	fmt.Printf("The temperature in Celsius on a sunlit Moon is %.2f and that is %.2f in Kelvin.\n", sunlitMoon, sunlitMoonKelvin)
}