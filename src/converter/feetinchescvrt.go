package converter

//ft2in converts feet to inches
func ft2in(f Feet) Inches {return Inches(f*12)}


//in2ft converts inches to feet
func in2ft(i Inches) Feet {return Feet(i/12)}

