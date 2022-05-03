package object

import (
	"math"

	"github.com/gernivisser/raytracer/engine/ray"
	"github.com/gernivisser/raytracer/tuples"
)

type Plane struct {
	*object
}

func NewPlane() *Plane {
	plane := Plane{
		object: NewObject(),
	}
	return &plane
}

func (p *Plane) Intersect(r *ray.Ray) Intersections {

	r2 := p.ray_to_ObjectSpace(r)

	var res Intersections

	if math.Abs(r2.Direction.Y) < p.Epsilon {
		return res
	}

	t := r2.Origin.Negate().Y / (r2.Direction.Y)
	res = *res.Insert(Intersection{T: t, Object: p})

	return res
}

func (p *Plane) NormalAt(world_point *tuples.Tuple) *tuples.Tuple {
	inv, _ := p.Transform.Invert()
	obj_norm := tuples.NewPoint(0, 1, 0)
	wrld_norm := inv.Transpose().MultiplyTuple(obj_norm)
	wrld_norm.T = 0
	return wrld_norm.Normalize()
}
