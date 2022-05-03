package light

import (
	"math"

	"github.com/gernivisser/raytracer/engine/materials"
	"github.com/gernivisser/raytracer/tuples"
	"github.com/gernivisser/raytracer/view"
)

type Point_Light struct {
	Position  tuples.Tuple
	Intensity view.Color
}

func NewPointLight(position tuples.Tuple, intensity view.Color) *Point_Light {
	return &Point_Light{Position: position, Intensity: intensity}
}

func WhitePointLight(position tuples.Tuple) *Point_Light {
	return NewPointLight(position, *view.NewColor(1, 1, 1))
}

func Lighting(point tuples.Tuple, light Point_Light, material materials.Material, eyev tuples.Tuple, normalv tuples.Tuple, inShadow bool) *view.Color {
	var diff, spec view.Color

	eff_color := material.Color
	lightv := light.Position.Sub(&point).Normalize()
	ambient := eff_color.Multiply(material.Ambient)
	light_dot_normal := lightv.Dot(normalv)

	if light_dot_normal < 0 || inShadow {
		diff = *view.NewColor(0, 0, 0)

		spec = *view.NewColor(0, 0, 0)
	} else {
		diff = *eff_color.Multiply(material.Diffuse).Multiply(light_dot_normal)

		reflectv := *lightv.Negate().Reflect(normalv)
		reflectv_dot_eye := reflectv.Dot(eyev)

		if reflectv_dot_eye <= 0 {
			spec = *view.NewColor(0, 0, 0)
		} else {
			factor := math.Pow(reflectv_dot_eye, material.Shininess)
			spec = *light.Intensity.Multiply(material.Specular * factor)
		}
	}

	return ambient.Add(&diff).Add(&spec)
}
