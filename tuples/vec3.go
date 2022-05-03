package tuples

func NewVec3(x float64, y float64, z float64) *Tuple {
	return NewTuple(x, y, z, 0)
}

func (tuple Tuple) IsVec3() bool {
	if tuple.T != 0 {
		return false
	}

	return true
}
