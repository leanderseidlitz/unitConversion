package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//CToK converts Celisius to Kelvin
func CToK(c Celsius) Kelvin {
	return Kelvin(ZeroKelvin + c)
}

//FToK converts Farenheit to Kelvin
func FToK(f Fahrenheit) Kelvin {
	return Kelvin((f + 459.67) * 5 / 9)
}

//KToC converts Kelvin to Celsius
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

//KToF converts Kelvin to Farenhiet
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(9/5*(k-273) + 32)
}
