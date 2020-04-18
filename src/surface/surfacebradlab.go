// Surface computes an SVG rendering of a 3-D surface function

package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size
	cells         = 50                 // number of grid cells
	xyrange       = 30.0                //axis ranges from negative val to positive val
	xyscale       = width / 4 / xyrange //pixel per x or y unit
	zscale        = height * -0.8        //pixels per z unit
	angle         = math.Pi/90          // angle of x, y axes 30º

)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30º), cos(30º)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg'"+
		"style='stroke: grey; fill: white; stroke-width:0.7'"+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j-1)
			cx, cy := corner(i, j+2)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}
func corner(i, j int) (float64, float64) {
	//find point (x,y)mat corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.6)
	y := xyrange * (float64(j)/cells - 0.5)

	//computer surface height z
	z := f(x, y)

	//project x,y,z isometrically onto 2D SVG CANVAS (sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/3 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from 0,0
	return math.Sin(r) / r
}
