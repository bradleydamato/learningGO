package converter

import "fmt"

type Feet float64
type Inches float64

func (f Feet) String() string {return fmt.Sprintf("%g FEET", f) }
func (i Inches) String() string {return fmt.Sprintf("%g INCHES", i) }
