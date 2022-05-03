package object

import (
	"github.com/gernivisser/raytracer/engine/materials"
	"github.com/gernivisser/raytracer/engine/ray"
	"github.com/gernivisser/raytracer/matrix"
	"github.com/gernivisser/raytracer/tuples"
)

type Object interface {
	GetMaterial() *materials.Material
	GetTransform() *matrix.Matrix
	SetMaterial(materials.Material)
	SetTransform(matrix.Matrix)
	Intersect(r *ray.Ray) Intersections
	ray_to_ObjectSpace(ray *ray.Ray) ray.Ray
	NormalAt(world_point *tuples.Tuple) *tuples.Tuple
}

type object struct {
	Material  materials.Material
	Transform matrix.Matrix
	Epsilon   float64
}

func NewObject() *object {
	return &object{
		Transform: *matrix.Identity4,
		Material:  materials.DefaultMaterial(),
		Epsilon:   0.0001,
	}
}

func (obj *object) GetMaterial() *materials.Material {
	return &obj.Material
}

func (obj *object) GetTransform() *matrix.Matrix {
	return &obj.Transform
}

func (obj *object) SetMaterial(material materials.Material) {
	obj.Material = material
}

func (obj *object) SetTransform(tranceform matrix.Matrix) {
	obj.Transform = tranceform
}

func (obj *object) ray_to_ObjectSpace(ray *ray.Ray) ray.Ray {
	// invert ray to keep sphere at world origin but stil be able to trancelate it
	inv, _ := obj.Transform.Invert()

	return ray.Transform(inv)
}
