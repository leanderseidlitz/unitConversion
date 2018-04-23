package lengthconv

import (
	"fmt"
)

type Feet float64
type Meter float64
type Inch float64
type Centimeter float64

func (f Feet) String() string {
	return fmt.Sprintf("%gF", f)
}

func (m Meter) String() string {
	return fmt.Sprintf("%gM", m)
}
func (i Inch) String() string {
	return fmt.Sprintf("%gin", i)
}

func (c Centimeter) String() string {
	return fmt.Sprintf("%gcm", c)
}
