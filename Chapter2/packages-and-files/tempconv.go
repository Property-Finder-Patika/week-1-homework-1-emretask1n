// Package tempconv performs Celsius and Fahrenheit conversions.

// Exercise 2.1
// Add types, constants, and functions to tempconv for processing temperatures in the
// Kelvin scale, where zero Kelvin is -273.15℃ and a difference of 1K has the same magnitude as 1℃.

package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

// Kelvin Temperature °K
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100

	AbsoluteZeroK Kelvin  = 0
	FreezingK     Kelvin  = 273.15
	BoilingK      Kelvin  = 373.15
	DeltaCK       Celsius = -273.15
	DeltaKC       Kelvin  = 273.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%.2f°K", k) }

//!-
