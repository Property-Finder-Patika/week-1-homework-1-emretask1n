// --Notes--

// The type of a variable or expression defines the characteristics of the values it may take on,
// such as their size(number of bits or number of elements, perhaps), how they are represented
// internally, the intrinsic operations that can be performed on them, and the methods associated with them

// Package tempconv performs Celsius and Fahrenheit temperature computations.

package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// This package defines two types, Celsius and Fahrenheit, for the two units of temperature.
// Even though both have the same underlying type, float64, they are not the same type

// an explicit type conversion like Celsius(t) or Fahrenheit(t) is required to convert from a
// float64. Celsius(t) and Fahrenheit(t) are conversions, not function calls

// the functions CToF and FToC convert between the two scales; they do
// return different values.

// fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
// boilingF := CToF(BoilingC)
// fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
// fmt.Printf("%g\n", boilingF-FreezingC) // compile error: type mismatch

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// Many types de clare a String method of this form because it controls how values of the type
// appear when printed as a string by the fmt package

//c := FToC(212.0)
//fmt.Println(c.String()) // "100°C"
//fmt.Printf("%v\n", c) // "100°C"; no need to call String explicitly
//fmt.Printf("%s\n", c) // "100°C"
//fmt.Println(c) // "100°C"
//fmt.Printf("%g\n", c) // "100"; does not call String
//fmt.Println(float64(c)) // "100"; does not call String
