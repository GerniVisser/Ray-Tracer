package materials

import (
	"github.com/gernivisser/raytracer/view"
)

type Material struct {
	Color            view.Color
	Ambient          float64
	Diffuse          float64
	Specular         float64
	Shininess        float64
	Reflection       float64
	Tranceparency    float64
	Refractive_Index float64
}

func NewMaterial(
	color view.Color,
	ambient float64,
	diffuse float64,
	specular float64,
	shininess float64,
	reflective float64,
	transparency float64,
	refractive_Index float64) Material {
	return Material{
		Color:            color,
		Ambient:          ambient,
		Diffuse:          diffuse,
		Specular:         specular,
		Shininess:        shininess,
		Reflection:       reflective,
		Tranceparency:    transparency,
		Refractive_Index: refractive_Index,
	}
}

func DefaultMaterial() Material {
	return NewMaterial(*view.NewColor(1, 1, 1), 0.1, 0.9, 1, 200, 0, 0, 1)
}

func TransparentMaterial() Material {
	return NewMaterial(*view.NewColor(1, 1, 1), 0.1, 0.1, 1, 20, 0, 1, 1.5)
}
