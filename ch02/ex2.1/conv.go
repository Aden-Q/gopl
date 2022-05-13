package tempconv

import "gopl.io/ch2/tempconv"

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK converts a Celsius temperature to Kelven.
func CToK(c Celsius) Kelvin { return Kelvin(c - Celsius(tempconv.AbsoluteZeroC)) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK converts a Fahrenheit temperature to Kelven.
func FToK(f Fahrenheit) Kelvin { return Kelvin(CToK(FToC(f))) }
