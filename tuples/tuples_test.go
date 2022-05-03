package tuples

import (
	"math"
	"testing"
)

func TestVec3(t *testing.T) {
	a := NewVec3(1, 1, 1)

	if !a.IsVec3() {
		t.Fatalf("a is not a vactor")
	}
}
func TestPoint(t *testing.T) {
	a := NewPoint(1, 1, 1)

	if !a.IsPoint() {
		t.Fatalf("a is not a point")
	}
}

func TestEqual(t *testing.T) {
	a := NewPoint(0.0001, 1, 1)
	b := NewPoint(0.00001, 1, 1)

	if !a.Equals(b) {
		t.Fatalf("values are not equal")
	}
}

func TestAdd(t *testing.T) {
	a := NewPoint(1, -1, 0)
	b := NewVec3(1, 0, 1)

	res := NewPoint(2, -1, 1)

	if !a.Add(b).Equals(res) {
		t.Fatalf("Addistion not working\nShould be ")
	}
}

func TestSub(t *testing.T) {
	a := NewPoint(1, -1, 0)
	b := NewPoint(1, 0, 1)

	add := a.Sub(b)
	res := NewPoint(0, -1, -1)

	if !add.Equals(res) {
		t.Fatalf("Subtractking not working")
	}
}

func TestNegation(t *testing.T) {
	a := NewPoint(1, -1, 0)

	add := a.Negate()
	res := NewPoint(-1, 1, 0)

	if !add.Equals(res) {
		t.Fatalf("Negation not working")
	}
}

func TestMultiplication(t *testing.T) {
	a := NewPoint(1, -2, 3)

	add := a.Multiply(3.5)
	res := NewPoint(3.5, -7, 10.5)

	if !add.Equals(res) {
		t.Fatalf("Multiplication not working")
	}
}

func TestDevision(t *testing.T) {
	a := NewPoint(1, -2, 3)

	add := a.Divide(2)
	res := NewPoint(0.5, -1, 1.5)

	if !add.Equals(res) {
		t.Fatalf("Division not working")
	}
}

func TestMagnitude(t *testing.T) {
	a := NewVec3(1, -2, 3)

	add := a.Magnitude()
	res := math.Sqrt(14)

	if add != res {
		t.Fatalf("Magnatude not working")
	}
}

func TestNormalize(t *testing.T) {
	a := NewVec3(4, 0, 0)

	x1 := a.Normalize()
	res := NewPoint(1, 0, 0)

	if !x1.Equals(res) {
		t.Fatalf("Normalizing not working")
	}
}

func TestDot(t *testing.T) {
	a := NewVec3(1, 2, 3)
	b := NewVec3(2, 3, 4)

	x1 := a.Dot(*b)
	res := 20.0

	if x1 != res {
		t.Fatalf("Dot not working")
	}
}

func TestCross(t *testing.T) {
	a := NewVec3(0, 1, 0)
	b := NewVec3(1, 0, 0)

	x1 := Cross(*b, *a)
	res := NewVec3(0, 0, -1)

	if x1.Equals(res) {
		t.Fatalf("Cross not working")
	}
}
