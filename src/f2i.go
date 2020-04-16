
//f2i converts feet to inches, pretty trivial problem but I'm learning so!

package main 

import (
"converter"
"fmt"
"os"
"strconv"
)

func main() {
	for _,arg := range os.Args[1:] {
		measurement, err := strconv.ParseFloat(arg,64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "f2i: %v\n", err)
			os.Exit(1)
		}
		f := converter.Feet(measurement)
		i := converter.Inches(measurement)
		fmt.Printf("%s = %s, %s = %s,\n",
			f,converter.Ft2in(f),i,converter.In2ft(i))
		}

	} 


