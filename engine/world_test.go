package engine

import (
	"math"
	"testing"

	"github.com/gernivisser/raytracer/engine/light"
	"github.com/gernivisser/raytracer/engine/materials"
	object "github.com/gernivisser/raytracer/engine/objects"
	"github.com/gernivisser/raytracer/engine/ray"
	"github.com/gernivisser/raytracer/matrix"
	"github.com/gernivisser/raytracer/tuples"
	"github.com/gernivisser/raytracer/view"
)

func TestWorld(t *testing.T) {
	var lights []light.Point_Light
	var objects []object.Object

	lights = append(lights, *light.WhitePointLight(*tuples.NewPoint(-10, 10, -10)))

	objects = append(objects, object.NewSphere())

	w := NewWorld(objects, lights)

	if len(w.Lights) != 1 {
		t.Fatalf("Error")
	}

}

func TestIntersecWorld(t *testing.T) {
	w := DefaultWorld()
	r := ray.NewRay(*tuples.NewPoint(0, 0, -5), *tuples.NewVec3(0, 0, 1))

	xs := w.Instersect_world(r)

	if len(xs) != 4 || xs[0].T != 4 || xs[2].T != 5.5 {
		t.Fatalf("Error")
	}
}

func TestPreCompute(t *testing.T) {
	s := object.NewPlane()
	r := ray.NewRay(*tuples.NewPoint(0, 1, -1), *tuples.NewVec3(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))

	i := object.Intersection{T: math.Sqrt(2), Object: s}
	intersections := object.NewIntersections(i)

	xs := i.Perpare_computation(*r, &intersections)

	if xs.Reflect_v != *tuples.NewVec3(0, math.Sqrt(2)/2, math.Sqrt(2)/2) {
		t.Fatalf("Error")
	}
}

func TestShadingInteraction(t *testing.T) {
	w := DefaultWorld()
	shape := object.NewPlane()

	shape.Material.Reflection = 0.5
	shape.Transform = *matrix.Translate(0, -1, 0)

	w.Objects = append(w.Objects, shape)
	r := ray.NewRay(*tuples.NewPoint(0, 0, -3), *tuples.NewVec3(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))

	i := object.Intersection{T: math.Sqrt(2), Object: shape}
	intersections := object.NewIntersections(i)

	xs := i.Perpare_computation(*r, &intersections)

	plight := w.Shade_Hit(*xs, 4)

	if !plight.Equals(view.NewColor(0.87677, 0.92436, 0.82918)) {
		t.Fatalf("Error")
	}
}

func TestShadingInteractionInside(t *testing.T) {
	w := DefaultWorld()
	w.Lights[0] = *light.WhitePointLight(*tuples.NewPoint(0, 0.25, 0))
	shape := w.Objects[1]
	r := ray.NewRay(*tuples.NewPoint(0, 0, 0), *tuples.NewVec3(0, 0, 1))

	i := object.Intersection{T: 0.5, Object: shape}
	intersections := object.NewIntersections(i)

	xs := i.Perpare_computation(*r, &intersections)

	plight := w.Shade_Hit(*xs, 4)

	if !plight.Equals(view.NewColor(0.90498, 0.90498, 0.90498)) {
		t.Fatalf("Error")
	}
}

func TestColorAt(t *testing.T) {
	w := DefaultWorld()
	outer := w.Objects[0].GetMaterial()
	outer.Ambient = 1
	w.Objects[0].SetMaterial(*outer)
	inner := w.Objects[1].GetMaterial()
	inner.Ambient = 1
	w.Objects[1].SetMaterial(*inner)
	r := ray.NewRay(*tuples.NewPoint(0, 0, 0.75), *tuples.NewVec3(0, 0, -1))
	c := w.Color_At(*r, 4)

	if !c.Equals(&inner.Color) {
		t.Errorf("Fail")
	}
}

func TestDefaultTransform(t *testing.T) {
	from := tuples.NewVec3(0, 0, 0)
	to := tuples.NewVec3(0, 0, 1)
	up := tuples.NewVec3(0, 1, 0)

	m := ViewTransform(*from, *to, *up)

	if !m.Equals(matrix.Scale(-1, 1, -1)) {
		t.Fatalf("Fail")
	}
}

func TestTranslateTransform(t *testing.T) {
	from := tuples.NewVec3(0, 0, 8)
	to := tuples.NewVec3(0, 0, 0)
	up := tuples.NewVec3(0, 1, 0)

	m := ViewTransform(*from, *to, *up)

	if !m.Equals(matrix.Translate(0, 0, -8)) {
		t.Fatalf("Fail")
	}
}

func TestTransformMatrix(t *testing.T) {
	from := tuples.NewVec3(1, 3, 2)
	to := tuples.NewVec3(4, -2, 8)
	up := tuples.NewVec3(1, 1, 0)

	m := ViewTransform(*from, *to, *up)

	m1 := matrix.PopulateMatrix([][]float64{
		{-0.50709, 0.50709, 0.67612, -2.36643},
		{0.76772, 0.60609, 0.12122, -2.82843},
		{-0.35857, 0.59761, -0.71714, 0.00000},
		{0, 0, 0, 1},
	})

	if !m.Equals(m1) {
		t.Errorf("Fail")
	}
}

func TestShadow1(t *testing.T) {
	w := DefaultWorld()
	p := tuples.NewPoint(0, 10, 0)
	sh := w.is_Shadowed(*p, w.Lights[0])

	if sh {
		t.Fatalf("Fail 1")
	}

	p = tuples.NewPoint(10, -10, 10)
	sh = w.is_Shadowed(*p, w.Lights[0])

	if !sh {
		t.Fatalf("Fail 2")
	}

	p = tuples.NewPoint(-20, 20, -20)
	sh = w.is_Shadowed(*p, w.Lights[0])

	if sh {
		t.Fatalf("Fail 3")
	}

	p = tuples.NewPoint(-2, 2, -2)
	sh = w.is_Shadowed(*p, w.Lights[0])

	if sh {
		t.Fatalf("Fail 4")
	}
}

func TestReflect_Color(t *testing.T) {
	w := DefaultWorld()
	r := ray.NewRay(*tuples.NewPoint(0, 0, 0), *tuples.NewVec3(0, 0, 1))

	shape := w.Objects[1]
	mat := shape.GetMaterial()
	mat.Ambient = 1
	shape.SetMaterial(*mat)

	i := object.Intersection{T: 1, Object: shape}
	intersections := object.NewIntersections(i)

	comps := i.Perpare_computation(*r, &intersections)
	col := w.Reflected_Color(comps, 4)

	if !col.Equals(view.NewColor(0, 0, 0)) {
		t.Fatalf("fail")
	}

}

func TestReflect_Color2(t *testing.T) {
	w := DefaultWorld()
	plane := object.NewPlane()

	plane.Material.Reflection = 0.5
	plane.Transform = *matrix.Translate(0, -1, 0)

	w.Objects = append(w.Objects, plane)

	r := ray.NewRay(*tuples.NewPoint(0, 0, -3), *tuples.NewVec3(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))

	i := object.Intersection{T: math.Sqrt(2), Object: plane}
	intersections := object.NewIntersections(i)

	comps := i.Perpare_computation(*r, &intersections)
	col := w.Reflected_Color(comps, 4)

	if !col.Equals(view.NewColor(0.19032, 0.2379, 0.14274)) {
		t.Fatalf("fail")
	}

}

func TestRefectionRercurrsion(t *testing.T) {
	w := DefaultWorld()
	//Floor
	plane1 := object.NewPlane()

	plane1.Material.Reflection = 1
	plane1.Transform = *matrix.Translate(0, -1, 0)

	w.Objects = append(w.Objects, plane1)
	// Ceiling
	plane2 := object.NewPlane()

	plane2.Material.Reflection = 1
	plane2.Transform = *matrix.Translate(0, 1, 0)

	w.Objects = append(w.Objects, plane2)

	r := ray.NewRay(*tuples.NewPoint(0, 0, 0), *tuples.NewVec3(0, 1, 0))

	col := w.Color_At(*r, 4)

	if !col.Equals(view.NewColor(0.1, 0.1, 0.1)) {
		t.Fatalf("fail")
	}

}

func TestRefractionIndex(t *testing.T) {
	s1 := object.NewSphere()
	s1.Transform = *matrix.Scale(2, 2, 2)
	s1.Material = materials.TransparentMaterial()
	s1.Material.Refractive_Index = 1.5

	s2 := object.NewSphere()
	s2.Transform = *matrix.Translate(0, 0, -0.25)
	s2.Material = materials.TransparentMaterial()
	s2.Material.Refractive_Index = 2

	s3 := object.NewSphere()
	s3.Transform = *matrix.Translate(0, 0, 0.25)
	s3.Material = materials.TransparentMaterial()
	s3.Material.Refractive_Index = 2.5

	r := ray.NewRay(*tuples.NewPoint(0, 0, -4), *tuples.NewVec3(0, 0, 1))

	xs := object.NewIntersections(
		object.Intersection{2, s1},
		object.Intersection{2.75, s2},
		object.Intersection{3.25, s3},
		object.Intersection{4.75, s2},
		object.Intersection{5.25, s3},
		object.Intersection{6, s1},
	)

	res := make([][]float64, 6)
	for i := 0; i < 6; i++ {
		res[i] = make([]float64, 2)
	}
	res[0][0] = 1
	res[0][1] = 1.5
	res[1][0] = 1.5
	res[1][1] = 2
	res[2][0] = 2
	res[2][1] = 2.5
	res[3][0] = 2.5
	res[3][1] = 2.5
	res[4][0] = 2.5
	res[4][1] = 1.5
	res[5][0] = 1.5
	res[5][1] = 1

	for i := 0; i < 6; i++ {
		comps := xs[i].Perpare_computation(*r, &xs)

		if comps.N1 != res[i][0] || comps.N2 != res[i][1] {
			t.Fatalf("Error %d", i)
		}
	}

}

func TestRefract_Color(t *testing.T) {
	w := DefaultWorld()
	r := ray.NewRay(*tuples.NewPoint(0, 0, -5), *tuples.NewVec3(0, 0, 1))

	shape := w.Objects[0]
	mat := *shape.GetMaterial()
	mat.Tranceparency = 1
	mat.Refractive_Index = 1.5
	shape.SetMaterial(mat)

	i := object.Intersection{T: 4, Object: shape}
	intersections := object.NewIntersections(i)

	comps := i.Perpare_computation(*r, &intersections)
	col := w.Refracted_Color(comps, 0)

	if !col.Equals(view.NewColor(0, 0, 0)) {
		t.Fatalf("fail")
	}

}

func TestTotalInternmalRefract_Color(t *testing.T) {
	w := DefaultWorld()
	r := ray.NewRay(*tuples.NewPoint(0, 0, math.Sqrt2/2), *tuples.NewVec3(0, 1, 0))

	shape := w.Objects[0]
	mat := *shape.GetMaterial()
	mat.Tranceparency = 1
	mat.Refractive_Index = 1.5
	shape.SetMaterial(mat)

	i1 := object.Intersection{T: math.Sqrt2 / 2, Object: shape}
	i2 := object.Intersection{T: -math.Sqrt2 / 2, Object: shape}
	intersections := object.NewIntersections(i1, i2)

	comps := i1.Perpare_computation(*r, &intersections)
	col := w.Refracted_Color(comps, 5)

	if !col.Equals(view.NewColor(0, 0, 0)) {
		t.Fatalf("fail")
	}
}

func TestRefract_Color2(t *testing.T) {
	w := DefaultWorld()

	floor := object.NewPlane()
	floor.Transform = *matrix.Translate(0, -1, 0)
	mat := *floor.GetMaterial()
	mat.Tranceparency = 0.5
	mat.Refractive_Index = 1.5
	floor.SetMaterial(mat)

	ball := object.NewSphere()
	ball.Transform = *matrix.Translate(0, -3.5, -0.5)
	mat = *ball.GetMaterial()
	mat.Color = *view.NewColor(1, 0, 0)
	mat.Ambient = 0.5
	ball.SetMaterial(mat)

	w.Objects = append(w.Objects, floor, ball)

	r := ray.NewRay(*tuples.NewPoint(0, 0, -3), *tuples.NewVec3(0, -math.Sqrt2/2, math.Sqrt2/2))

	i1 := object.Intersection{T: math.Sqrt2, Object: floor}
	intersections := object.NewIntersections(i1)

	comps := i1.Perpare_computation(*r, &intersections)
	col := w.Shade_Hit(*comps, 5)

	if !col.Equals(view.NewColor(0.93642, 0.68642, 0.68642)) {
		t.Fatalf("fail")
	}
}
