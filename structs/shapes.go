package shapes

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (r Circle) Area() float64 {
	return r.Radius * r.Radius * math.Pi
}

type Triangle struct {
	Base   float64
	Height float64
}

func (r Triangle) Area() float64 {
	return (r.Base * r.Height) / 2
}
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}
