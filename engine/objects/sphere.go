package object

import (
	"math"

	"github.com/gernivisser/raytracer/engine/ray"
	"github.com/gernivisser/raytracer/tuples"
)

type Sphere struct {
	*object
}

func NewSphere() *Sphere {
	sphere := Sphere{
		object: NewObject(),
	}
	return &sphere
}

func (s *Sphere) Intersect(r *ray.Ray) Intersections {

	r2 := s.ray_to_ObjectSpace(r)
	//Calculating discriminat
	sphere_to_ray := r2.Origin.Sub(tuples.NewPoint(0, 0, 0))

	a := r2.Direction.Dot(r2.Direction)
	b := r2.Direction.Dot(*sphere_to_ray) * 2
	c := sphere_to_ray.Dot(*sphere_to_ray) - 1

	discriminant := math.Pow(b, 2) - (4 * a * c)

	var res Intersections

	// if discriminat < 0  then the ray does not intersect the sphere
	if discriminant < 0 {
		return res
	}

	res = *res.Insert(Intersection{T: (-b - math.Sqrt(discriminant)) / (2 * a), Object: s})
	res = *res.Insert(Intersection{T: (-b + math.Sqrt(discriminant)) / (2 * a), Object: s})

	return res
}

func (s *Sphere) NormalAt(world_point *tuples.Tuple) *tuples.Tuple {
	inv, _ := s.Transform.Invert()
	obj_point := inv.MultiplyTuple(world_point)
	obj_norm := obj_point.Sub(tuples.NewPoint(0, 0, 0))
	wrld_norm := inv.Transpose().MultiplyTuple(obj_norm)
	wrld_norm.T = 0
	return wrld_norm.Normalize()
}
