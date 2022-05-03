package tuples

func NewPoint(x float64, y float64, z float64) *Tuple {
	return NewTuple(x, y, z, 1)
}

func (tuple Tuple) IsPoint() bool {
	if tuple.T != 1 {
		return false
	}

	return true
}
