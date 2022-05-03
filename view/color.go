package view

import (
	"fmt"
	"math"
)

type Color struct {
	Red   float64
	Green float64
	Blue  float64
}

func NewColor(red float64, green float64, blue float64) *Color {
	return &Color{Red: red, Green: green, Blue: blue}
}

func (c *Color) Print() {
	s := "Color \n"
	s += fmt.Sprintf("Red: %f\n", c.Red)
	s += fmt.Sprintf("Green: %f\n", c.Green)
	s += fmt.Sprintf("Blue: %f\n", c.Blue)

	fmt.Print(s)
}

func (c *Color) Add(c2 *Color) *Color {
	return NewColor(c.Red+c2.Red, c.Green+c2.Green, c.Blue+c2.Blue)
}

// Subtract returns the difference between two colors.
func (c *Color) Subtract(c2 *Color) *Color {
	return NewColor(c.Red-c2.Red, c.Green-c2.Green, c.Blue-c2.Blue)
}

// Multiply returns a color multiplied by a value.
func (c *Color) Multiply(val float64) *Color {
	return NewColor(c.Red*val, c.Green*val, c.Blue*val)
}

// MultiplyColor multiplies two colors by each other and returns the result.
func (c *Color) MultiplyColor(c2 *Color) *Color {
	return NewColor(c.Red*c2.Red, c.Green*c2.Green, c.Blue*c2.Blue)
}

func (c *Color) Equals(c2 *Color) bool {
	var epsilon float64 = 0.0001

	X := math.Abs(c.Red - c2.Red)
	Y := math.Abs(c.Green - c2.Green)
	Z := math.Abs(c.Blue - c2.Blue)

	if X > epsilon || Y > epsilon || Z > epsilon {
		return false
	}
	return true
}
