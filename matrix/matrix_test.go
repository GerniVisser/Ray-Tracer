package matrix

import (
	"testing"

	"github.com/gernivisser/raytracer/tuples"
)

func TestEqual(t *testing.T) {
	a := NewMatrix(4)

	for x := range a.data {
		for y := range a.data {
			a.data[x][y] = 3
		}
	}

	b := NewMatrix(4)
	for x := range b.data {
		for y := range b.data {
			b.data[x][y] = 3.00
		}
	}
	if !a.Equals(b) {
		t.Fatalf("Matrix A and B are not equal")
	}
}

func TestMuti(t *testing.T) {
	a1 := [][]float64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 8, 7, 6}, {5, 4, 3, 2}}
	a := PopulateMatrix(a1)

	b1 := [][]float64{{-2, 1, 2, 3}, {3, 2, 1, -1}, {4, 3, 6, 5}, {1, 2, 7, 8}}
	b := PopulateMatrix(b1)

	res1 := [][]float64{{20, 22, 50, 48}, {44, 54, 114, 108}, {40, 58, 110, 102}, {16, 26, 46, 42}}
	res := PopulateMatrix(res1)

	mult := a.Multiply(b)

	if !mult.Equals(res) {
		t.Fatalf("Failed")
	}
}

func TestMutiTuple(t *testing.T) {
	a1 := [][]float64{{1, 2, 3, 4}, {2, 4, 4, 2}, {8, 6, 4, 1}, {0, 0, 0, 1}}
	a := PopulateMatrix(a1)
	b1 := tuples.NewPoint(1, 2, 3)
	res1 := tuples.NewTuple(18, 24, 33, 1)
	mult := a.MultiplyTuple(b1)

	val := mult.Equals(res1)

	if !val {
		t.Fatalf("Error ")
	}
}

func TestIdentity(t *testing.T) {
	a1 := [][]float64{{1, 2, 3, 4}, {2, 4, 4, 2}, {8, 6, 4, 1}, {0, 0, 0, 1}}
	a := PopulateMatrix(a1)

	res := Identity4.Multiply(a)

	if !res.Equals(a) {
		t.Fatalf("Failed")
	}
}

func TestTranspose(t *testing.T) {
	a1 := [][]float64{{1, 2, 3, 4}, {2, 4, 4, 2}, {8, 6, 4, 1}, {0, 0, 0, 1}}
	a := PopulateMatrix(a1)

	b1 := [][]float64{{1, 2, 8, 0}, {2, 4, 6, 0}, {3, 4, 4, 0}, {4, 2, 1, 1}}
	b := PopulateMatrix(b1)

	res := a.Transpose()

	if !res.Equals(b) {
		t.Fatalf("Failed")
	}
}

func TestDeterminant(t *testing.T) {
	a1 := [][]float64{{1, 2}, {2, 3}}
	a := PopulateMatrix(a1)

	b1 := float64(-1)

	res := a.Determinant()

	if res != b1 {
		t.Fatalf("error")
	}
}

func TestInverseDeep(t *testing.T) {
	a1 := [][]float64{{8, -5, 9, 2}, {7, 5, 6, 1}, {-6, 0, 9, 6}, {-3, 0, -9, -4}}
	a := PopulateMatrix(a1)

	b1 := [][]float64{
		{-0.15385, -0.15385, -0.28205, -0.53846},
		{-0.07692, 0.12308, 0.02564, 0.03077},
		{0.35897, 0.35897, 0.43590, 0.92308},
		{-0.69231, -0.69231, -0.76923, -1.92308}}
	b := PopulateMatrix(b1)

	res, _ := a.Invert()
	if !res.Equals(b) {
		t.Fatalf("failed")
	}
}

func TestInverseMutl(t *testing.T) {
	a1 := [][]float64{{3, -9, 7, 3}, {3, -8, 2, -9}, {-4, 4, 4, 1}, {-6, 5, -1, 1}}
	a := PopulateMatrix(a1)

	b1 := [][]float64{{8, 2, 2, 2}, {3, -1, 7, 0}, {7, 0, 5, 4}, {6, -2, 0, 5}}
	b := PopulateMatrix(b1)

	mult := a.Multiply(b)

	inv, _ := b.Invert()

	x := mult.Multiply(inv)

	if !x.Equals(a) {
		t.Fatalf("Test Failed")
	}
}

func TestTranselation(t *testing.T) {
	trans := Translate(5, -3, 2)
	inv, _ := trans.Invert()
	p := tuples.NewPoint(-3, 4, 5)
	res := inv.MultiplyTuple(p)

	awns := tuples.NewPoint(-8, 7, 3)

	if !awns.Equals(res) {
		t.Fatalf("failed")
	}
}

func TestTranselationVector(t *testing.T) {
	trans := Translate(5, -3, 2)
	p := tuples.NewVec3(-3, 4, 5)
	res := trans.MultiplyTuple(p)

	if !p.Equals(res) {
		t.Fatalf("failed")
	}
}
