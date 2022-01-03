package structsmethodsinterfaces

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (rectangle Rectangle) Perimeter() (perimeter float64) {
	return math.Round(2*(rectangle.Width+rectangle.Height)*100) / 100
}

func (rectangle Rectangle) Area() (area float64) {
	return math.Round((rectangle.Width*rectangle.Height)*100) / 100
}

func (circle Circle) Area() (area float64) {
	return math.Round(math.Pi*circle.Radius*circle.Radius*100) / 100
}

func (triangle Triangle) Area() (area float64) {
	return math.Round((triangle.Height*triangle.Base)/2*100) / 100
}
