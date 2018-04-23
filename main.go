package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/midjul/unitConversion/lengthconv"
	"github.com/midjul/unitConversion/tempconv"
)

var cf = flag.Float64("ctf", 0, "convert celsius to farenheit")
var fc = flag.Float64("ftc", 0, "convert farenheit to celsius")
var ck = flag.Float64("ctk", 0, "convert celsius to kelvin")
var fk = flag.Float64("ktc", 0, "convert kelvin to celsius")

var mf = flag.Float64("mtf", 0, "convet meter to feet")
var fm = flag.Float64("ftm", 0, "convet feet to meter")
var ic = flag.Float64("itc", 0, "convert incehst to centimeters")
var ci = flag.Float64("cti", 0, "convert centimeters to inches")

func main() {
	flag.Parse()

	if cf != nil {
		res := tempconv.CToF(tempconv.Celsius(*cf))
		fmt.Fprintf(os.Stdout, "%vC is %v\n", *cf, res)
	}

	if fc != nil {
		res := tempconv.FToC(tempconv.Fahrenheit(*cf))
		fmt.Fprintf(os.Stdout, "%vF is %v\n", *cf, res)

	}
	if ck != nil {
		res := tempconv.CToK(tempconv.Celsius(*cf))
		fmt.Fprintf(os.Stdout, "%vC is %v\n", *cf, res)

	}
	if fk != nil {
		res := tempconv.FToK(tempconv.Fahrenheit(*cf))
		fmt.Fprintf(os.Stdout, "%vF is %v\n", *cf, res)

	}
	if mf != nil {
		res := lengthconv.MToF(lengthconv.Meter(*mf))
		fmt.Fprintf(os.Stdout, "%vM is %v\n", *mf, res)

	}

	if fm != nil {
		res := lengthconv.FToM(lengthconv.Feet(*fm))
		fmt.Fprintf(os.Stdout, "%vFeet is %v\n", *fm, res)

	}
	if ic != nil {
		res := lengthconv.IToC(lengthconv.Inch(*ic))
		fmt.Fprintf(os.Stdout, "%vin is %v\n", *ic, res)

	}
	if ci != nil {
		res := lengthconv.CToI(lengthconv.Centimeter(*ci))
		fmt.Fprintf(os.Stdout, "%vcm is %v\n", *ci, res)

	}

}
