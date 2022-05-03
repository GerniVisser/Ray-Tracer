package ray

import (
	"github.com/gernivisser/raytracer/matrix"
	"github.com/gernivisser/raytracer/tuples"
)

type Ray struct {
	Origin    tuples.Tuple
	Direction tuples.Tuple
}

func NewRay(origin tuples.Tuple, direction tuples.Tuple) *Ray {
	ray := Ray{Origin: origin, Direction: direction}
	return &ray
}

func (r *Ray) Position(t float64) *tuples.Tuple {
	vec := r.Direction.Multiply(t)
	res := r.Origin.Add(vec)

	return res
}

func (r *Ray) Transform(transformation *matrix.Matrix) Ray {

	direction := transformation.MultiplyTuple(&r.Direction)
	position := transformation.MultiplyTuple(&r.Origin)

	return Ray{Origin: *position, Direction: *direction}
}
