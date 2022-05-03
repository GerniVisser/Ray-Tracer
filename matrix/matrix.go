package matrix

import (
	"fmt"
	"math"

	"github.com/gernivisser/raytracer/tuples"
)

type Matrix struct {
	data [][]float64
	size int
}

// Epsilon is used when testing for equility to neglect small differences that can occure with floats
var epsilon float64 = 0.0001

// Identity Matrix
var identity4data = [][]float64{
	{1, 0, 0, 0},
	{0, 1, 0, 0},
	{0, 0, 1, 0},
	{0, 0, 0, 1},
}
var Identity4 = PopulateMatrix(identity4data)

func NewMatrix(size int) *Matrix {
	data := make([][]float64, size)
	for i := range data {
		data[i] = make([]float64, size)
	}
	m := Matrix{data: data, size: size}
	return &m
}

func PopulateMatrix(data [][]float64) *Matrix {
	m := Matrix{data: data, size: len(data)}
	return &m
}

// Utility function to log contents of matix to console
func (m *Matrix) Print() {
	s := fmt.Sprintf("%d", m.size) + " x " + fmt.Sprintf("%d", m.size) + " Matrix\n"
	for row := range m.data {
		for col := range m.data {
			s += fmt.Sprintf("%f", m.data[row][col]) + "\t| "
		}
		s += "\n"
	}
	fmt.Print(s)
}

func (m *Matrix) CompareSize(m2 *Matrix) bool {
	if m.size != m2.size {
		panic("Matrices must be the same size")
	}
	if (m == nil) != (m2 == nil) {
		panic("Matrices cannot be nil")
	}
	return true

}

func (m *Matrix) Equals(m2 *Matrix) bool {
	m.CompareSize(m2)

	for row := range m.data {
		for col := range m.data {
			if math.Abs(m.data[row][col]-m2.data[row][col]) >= epsilon {
				return false
			}
		}
	}

	return true
}

func (m *Matrix) Multiply(m2 *Matrix) *Matrix {
	m.CompareSize(m2)

	res := NewMatrix(m.size)

	for row := range m.data {
		for col := range m.data {
			res.data[row][col] = m.data[row][0]*m2.data[0][col] +
				m.data[row][1]*m2.data[1][col] +
				m.data[row][2]*m2.data[2][col] +
				m.data[row][3]*m2.data[3][col]
		}
	}

	return res
}

func (m *Matrix) MultiplyTuple(t *tuples.Tuple) *tuples.Tuple {
	if m.size != 4 {
		panic("Matrix and Tuple does not have same dimentions must be a 4x4 matrix")
	}

	var res [4]float64
	for i := 0; i <= 3; i++ {
		res[i] = m.data[i][0]*t.X +
			m.data[i][1]*t.Y +
			m.data[i][2]*t.Z +
			m.data[i][3]*float64(t.T)
	}

	return tuples.NewTuple(res[0], res[1], res[2], int8(res[3]))

}

// Flips the rows and the coloums
func (m *Matrix) Transpose() *Matrix {
	res := NewMatrix(m.size)
	for row := range m.data {
		for col := range m.data {
			res.data[col][row] = m.data[row][col]
		}
	}

	return res
}

func (m *Matrix) Determinant() float64 {
	var res float64 = 0

	if m.size == 2 {
		res = m.data[0][0]*m.data[1][1] - m.data[0][1]*m.data[1][0]
		return res
	} else {
		for i := range m.data {
			res += m.data[0][i] * m.cofactor(0, i)
		}
	}
	return res

}

func (m *Matrix) submatrix(row int, col int) *Matrix {
	res := NewMatrix(m.size - 1)
	var xh, yh int = 0, 0

	for x := 0; x <= m.size-2; x++ {
		for y := 0; y <= m.size-2; y++ {
			if x >= row {
				xh = 1
			}
			if y >= col {
				yh = 1
			}
			res.data[x][y] = m.data[x+xh][y+yh]
			xh, yh = 0, 0
		}
	}

	return res
}

func (m *Matrix) minor(row int, col int) float64 {
	sub := m.submatrix(row, col)
	res := sub.Determinant()

	return res
}

func (m *Matrix) cofactor(row int, col int) float64 {
	res := m.minor(row, col)

	if (row+col)%2 != 0 {
		return -res
	}
	return res
}

func (m *Matrix) Invertable() bool {
	if m.Determinant() == 0 {
		return false
	}
	return true
}

func (m *Matrix) Invert() (*Matrix, bool) {
	if !m.Invertable() {
		return nil, false
	}

	res := NewMatrix(m.size)

	for row := range m.data {
		for col := range m.data {
			co := m.cofactor(row, col)
			det := m.Determinant()

			res.data[col][row] = co / det
		}
	}

	return res, true
}

func Translate(x, y, z float64) *Matrix {
	data := [][]float64{
		{1, 0, 0, x},
		{0, 1, 0, y},
		{0, 0, 1, z},
		{0, 0, 0, 1},
	}

	return PopulateMatrix(data)
}

func Scale(x, y, z float64) *Matrix {
	data := [][]float64{
		{x, 0, 0, 0},
		{0, y, 0, 0},
		{0, 0, z, 0},
		{0, 0, 0, 1},
	}

	return PopulateMatrix(data)
}

func RotateX(radians float64) *Matrix {
	data := [][]float64{
		{1, 0, 0, 0},
		{0, math.Cos(radians), -math.Sin(radians), 0},
		{0, math.Sin(radians), math.Cos(radians), 0},
		{0, 0, 0, 1},
	}

	return PopulateMatrix(data)
}

func RotateY(radians float64) *Matrix {
	data := [][]float64{
		{math.Cos(radians), 0, math.Sin(radians), 0},
		{0, 1, 0, 0},
		{-math.Sin(radians), 0, math.Cos(radians), 0},
		{0, 0, 0, 1},
	}

	return PopulateMatrix(data)
}

func RotateZ(radians float64) *Matrix {
	data := [][]float64{
		{math.Cos(radians), -math.Sin(radians), 0, 0},
		{math.Sin(radians), math.Cos(radians), 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}

	return PopulateMatrix(data)
}

func Shear(xy, xz, yx, yz, zx, zy float64) *Matrix {
	data := [][]float64{
		{1, xy, xz, 0},
		{yx, 1, yz, 0},
		{zx, zy, 1, 0},
		{0, 0, 0, 1},
	}

	return PopulateMatrix(data)
}
