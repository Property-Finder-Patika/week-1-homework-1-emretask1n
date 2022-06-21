// Ftoc prints two Fahrenheit-to-Celsius conversions.

// --Notes--

//The function fToC below encapsulates the temperature conversion logic so that it is defined only once but may be used from
//multiple places. Here main calls it twice, using the values of two different local constants:
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"
}
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
