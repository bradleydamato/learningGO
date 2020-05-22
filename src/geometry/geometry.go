package geometry

import "math"

type Point struct{X,Y float64}

// journey connecting two points in a straight line 
type Path []Point

//traditional function
func Distance(p,q Point) float64 {
	return math.Hypot(q.X-p.X,q.Y-p.Y)
}

//same as above but as a method of the point type 
func (p Point) Distance(q point) float64{
	return math.Hypot(q.X-p.X,q.Y-p.Y)
}

func (path Path) Distance() float64{
	sum := 0.0 
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}