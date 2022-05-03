package engine

import (
	"math"

	"github.com/gernivisser/raytracer/engine/light"
	"github.com/gernivisser/raytracer/engine/materials"
	object "github.com/gernivisser/raytracer/engine/objects"
	"github.com/gernivisser/raytracer/engine/ray"
	"github.com/gernivisser/raytracer/matrix"
	"github.com/gernivisser/raytracer/tuples"
	"github.com/gernivisser/raytracer/view"
)

type World struct {
	Objects []object.Object
	Lights  []light.Point_Light
}

func NewWorld(objects []object.Object, lights []light.Point_Light) *World {
	return &World{Objects: objects, Lights: lights}
}

func DefaultWorld() *World {
	var lights []light.Point_Light
	var objects []object.Object

	lights = append(lights, *light.WhitePointLight(*tuples.NewPoint(-10, 10, -10)))

	obj1 := object.NewSphere()
	obj1.SetMaterial(materials.NewMaterial(*view.NewColor(0.8, 1, 0.6), 0.1, 0.7, 0.2, 200, 0, 0, 1))
	obj2 := object.NewSphere()
	obj2.SetTransform(*matrix.Scale(0.5, 0.5, 0.5))

	objects = append(objects, obj1, obj2)

	return NewWorld(objects, lights)
}

func (w *World) Instersect_world(r *ray.Ray) object.Intersections {
	var intersections object.Intersections

	for _, object := range w.Objects {
		ints := object.Intersect(r)
		for _, intersect := range ints {
			intersections = *intersections.Insert(intersect)
		}
	}

	return intersections
}

func (w *World) is_Shadowed(point tuples.Tuple, ligth light.Point_Light) bool {
	v := ligth.Position.Sub(&point)
	direction := v.Normalize()
	distance := v.Magnitude()
	r := ray.NewRay(point, *direction)

	intersections := w.Instersect_world(r)
	h := intersections.Hit()

	if h.Object != nil && h.T < distance {
		return true
	}

	return false
}

func (w *World) Reflected_Color(comps *object.IntersectComp, remaining int) *view.Color {
	if comps.Object.GetMaterial().Ambient == 1 || remaining == 0 {
		return view.NewColor(0, 0, 0)
	}

	r := ray.NewRay(comps.Over_point, comps.Reflect_v)

	return w.Color_At(*r, remaining-1).Multiply(comps.Object.GetMaterial().Reflection)
}

func (w *World) Refracted_Color(comps *object.IntersectComp, remaining int) *view.Color {
	if comps.Object.GetMaterial().Tranceparency == 0 || remaining == 0 {
		return view.NewColor(0, 0, 0)
	}

	// Compute angle of exit ray
	n_ratio := comps.N1 / comps.N2
	cos_i := comps.Eyev.Dot(comps.Normalv)
	sin2_t := n_ratio * n_ratio * (1 - cos_i*cos_i)

	// If sin2 is greated than 1 Total internal refraction occered
	if sin2_t > 1 {
		return view.NewColor(0, 0, 0)
	}

	// Compute the refracted ray
	cos_t := math.Sqrt(1.0 - sin2_t)

	directtion := comps.Normalv.Multiply(n_ratio*cos_i - cos_t).Sub(comps.Eyev.Multiply(n_ratio))

	r := ray.NewRay(comps.Under_point, *directtion)

	return w.Color_At(*r, remaining-1).Multiply(comps.Object.GetMaterial().Tranceparency)
}

func (w *World) Shade_Hit(comp object.IntersectComp, remaining int) *view.Color {
	var col view.Color = *view.NewColor(0, 0, 0)

	for _, l := range w.Lights {
		in_shade := w.is_Shadowed(comp.Over_point, l)
		col2 := light.Lighting(comp.Over_point, l, *comp.Object.GetMaterial(), comp.Eyev, comp.Normalv, in_shade)
		col = *col.Add(col2)
		reflected_col := w.Reflected_Color(&comp, remaining)
		col = *col.Add(reflected_col)
		refracted_col := w.Refracted_Color(&comp, remaining)
		col = *col.Add(refracted_col)
	}
	return &col
}

func (w *World) Color_At(r ray.Ray, remaining int) *view.Color {
	var col view.Color = *view.NewColor(0, 0, 0)

	i := w.Instersect_world(&r)
	if len(i) == 0 {
		return &col
	}
	for count := 0; count < len(i); count++ {
		if i[count].T > 0 {
			p := i[count].Perpare_computation(r, &i)
			col = *w.Shade_Hit(*p, remaining)

			return &col
		}
	}

	return &col
}

func ViewTransform(from tuples.Tuple, to tuples.Tuple, up tuples.Tuple) matrix.Matrix {
	forward := to.Sub(&from).Normalize()
	left := tuples.Cross(*forward, *up.Normalize())
	true_up := tuples.Cross(*left, *forward)

	orientation := matrix.PopulateMatrix([][]float64{
		{left.X, left.Y, left.Z, 0},
		{true_up.X, true_up.Y, true_up.Z, 0},
		{-forward.X, -forward.Y, -forward.Z, 0},
		{0, 0, 0, 1},
	})

	orientation = orientation.Multiply(matrix.Translate(-from.X, -from.Y, -from.Z))

	return *orientation
}
