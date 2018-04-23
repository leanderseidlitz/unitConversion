package lengthconv

//MToF converts Meter to Feet
func MToF(m Meter) Feet {
	return Feet(m * 3.2808)
}

//FToM converts Feet to Meter
func FToM(f Feet) Meter {
	return Meter(f / 3.2808)
}

//IToC convertes Inches to Centimeters
func IToC(i Inch) Centimeter {
	return Centimeter(i / 0.39370)
}

//CToI converts Centimeters to Inches
func CToI(c Centimeter) Inch {
	return Inch((c * 0.3937007874) / c)
}
