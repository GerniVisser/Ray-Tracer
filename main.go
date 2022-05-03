package main

import (
	"math"

	"github.com/gernivisser/raytracer/engine"
	"github.com/gernivisser/raytracer/engine/light"
	"github.com/gernivisser/raytracer/engine/materials"
	object "github.com/gernivisser/raytracer/engine/objects"
	"github.com/gernivisser/raytracer/matrix"
	"github.com/gernivisser/raytracer/tuples"
	"github.com/gernivisser/raytracer/view"
)

func main() {
	floor := object.NewPlane()
	floor.Material.Color = *view.NewColor(1, 0.9, 0.95)
	floor.Material.Specular = 0
	floor.Material.Ambient = 0.0
	floor.Material.Shininess = 20
	floor.Material.Reflection = 0.1

	left_wall := object.NewPlane()
	left_wall.Transform = *matrix.RotateX(math.Pi / 2).Multiply(matrix.Translate(0, 0, -5))
	left_wall.Material = floor.Material
	left_wall.Material.Reflection = 0
	left_wall.Material.Color = *view.NewColor(0.7, 0.7, 0.7)

	Right_wall := object.NewPlane()
	Right_wall.Transform = *matrix.RotateZ(math.Pi / 2).Multiply(matrix.Translate(-5, 0, 0))
	Right_wall.Material = floor.Material
	Right_wall.Material.Ambient = 0.4
	Right_wall.Material.Reflection = 0
	Right_wall.Material.Color = *view.NewColor(0.5, 0.5, 0.6)

	ball := object.NewSphere()
	ball.Material.Color = *view.NewColor(1, 0, 0)
	ball.Material.Shininess = 150
	ball.Material.Specular = 1
	ball.Transform = *matrix.Translate(-3, 0.5, -4.5).Multiply(matrix.Scale(0.5, 0.5, 0.5))

	ball2 := object.NewSphere()
	ball2.Material.Color = *view.NewColor(0.8, 0.8, 0.5)
	ball2.Material.Specular = 0.1
	ball2.Material.Ambient = 0
	ball2.Material.Reflection = 0.96
	ball2.Transform = *matrix.Translate(-4.5, 1, -2.5)

	ball3 := object.NewSphere()
	ball3.Material = materials.TransparentMaterial()
	ball3.Transform = *matrix.Translate(-5.5, 0.5, -5.5).Multiply(matrix.Scale(0.5, 0.5, 0.5))

	ball4 := object.NewSphere()
	ball4.Material = materials.DefaultMaterial()
	ball4.Material.Color = *view.NewColor(0.3, 0.7, 0.3)
	ball4.Material.Specular = 0.7
	ball4.Material.Shininess = 20
	ball4.Transform = *matrix.Translate(-6, 0.3, -4.8).Multiply(matrix.Scale(0.3, 0.3, 0.3))

	l := light.NewPointLight(*tuples.NewPoint(-14, 7, -8), *view.NewColor(0.7, 0.7, 0.7))

	w := engine.NewWorld([]object.Object{floor, left_wall, Right_wall, ball, ball2, ball3, ball4}, []light.Point_Light{*l})

	c := engine.NewCamera(1080, 720, math.Pi/3)
	c.Transform = engine.ViewTransform(*tuples.NewPoint(-7, 1, -7), *tuples.NewPoint(0, -1, 0), *tuples.NewVec3(0, 1, 0))

	canvas := c.Render(*w)
	canvas.WriteFile("Refraction")
}
