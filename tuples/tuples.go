package tuples

import (
	"fmt"
	"math"
)

type Tuple struct {
	X float64
	Y float64
	Z float64
	T int8
}

var epsilon float64 = 0.0001

func NewTuple(X float64, Y float64, Z float64, T int8) *Tuple {
	return &Tuple{X, Y, Z, T}
}

func (a Tuple) Equals(b *Tuple) bool {
	X := math.Abs(a.X - b.X)
	Y := math.Abs(a.Y - b.Y)
	Z := math.Abs(a.Z - b.Z)

	if X > epsilon || Y > epsilon || Z > epsilon {
		return false
	}
	return true
}

func (a Tuple) Print() {
	s := "Tuple\n"
	s += fmt.Sprintf("X: %f\n", a.X)
	s += fmt.Sprintf("Y: %f\n", a.Y)
	s += fmt.Sprintf("Z: %f\n", a.Z)

	fmt.Print(s)
}

func (a Tuple) Add(b *Tuple) *Tuple {
	X := NewTuple(a.X+b.X, a.Y+b.Y, a.Z+b.Z, a.T+b.T)

	return X
}

func (a Tuple) Sub(b *Tuple) *Tuple {
	X := NewTuple(a.X-b.X, a.Y-b.Y, a.Z-b.Z, a.T-b.T)

	return X
}

func (a Tuple) Negate() *Tuple {
	X := NewTuple(0-a.X, 0-a.Y, 0-a.Z, a.T)

	return X
}

func (a Tuple) Multiply(val float64) *Tuple {
	X := NewTuple(val*a.X, val*a.Y, val*a.Z, a.T)

	return X
}

func (a Tuple) Divide(val float64) *Tuple {
	X := NewTuple(a.X/val, a.Y/val, a.Z/val, a.T)

	return X
}

func (a Tuple) Magnitude() float64 {
	//Check type / Error handling
	if !a.IsVec3() {
		panic("Tuple must be of tyype Vec3 to obtain Magnatude")
	}

	determinant := math.Pow(a.X, 2) + math.Pow(a.Y, 2) + math.Pow(a.Z, 2)
	return math.Sqrt(determinant)
}

func (a Tuple) Normalize() *Tuple {
	if !a.IsVec3() {
		panic("Tuple must be of tyype Vec3 to Normalize")
	}

	mag := a.Magnitude()
	X := a.Divide(mag)
	return X
}

func (a Tuple) Dot(b Tuple) float64 {
	if !a.IsVec3() {
		panic("Tuple must be of tyype Vec3 to obtain Dot Product")
	}

	return (a.X*b.X + a.Y*b.Y + a.Z*b.Z)
}

func Cross(a Tuple, b Tuple) *Tuple {
	if !a.IsVec3() || !b.IsVec3() {
		panic("Both tuples have to be of type Tuple")
	}

	X := NewVec3(
		a.Y*b.Z-a.Z*b.Y,
		a.Z*b.X-a.X*b.Z,
		a.X*b.Y-a.Y*b.X)

	return X
}

func (a Tuple) Reflect(normal Tuple) *Tuple {
	x := a.Dot(normal) * 2
	return a.Sub(normal.Multiply(x))
}
