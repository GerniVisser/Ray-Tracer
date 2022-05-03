package object

import (
	"testing"

	"github.com/gernivisser/raytracer/engine/light"
	"github.com/gernivisser/raytracer/engine/materials"
	"github.com/gernivisser/raytracer/engine/ray"
	"github.com/gernivisser/raytracer/matrix"
	"github.com/gernivisser/raytracer/tuples"
	"github.com/gernivisser/raytracer/view"
)

func TestIntersectionMiddle(t *testing.T) {
	ray1 := ray.NewRay(*tuples.NewPoint(0, 0, -5), *tuples.NewVec3(0, 0, 1))

	sphere := NewSphere()

	var res Intersections = sphere.Intersect(ray1)
	var awnsr [2]float64

	awnsr[0] = 4
	awnsr[1] = 6

	if res[0].T != (awnsr[0]) || res[1].T != (awnsr[1]) {
		t.Fatalf("Middle test failed")
	}

}

func TestIntersectionTangent(t *testing.T) {
	ray1 := ray.NewRay(*tuples.NewPoint(-5, 1, 0), *tuples.NewVec3(1, 0, 0))

	sphere := NewSphere()

	var res Intersections = sphere.Intersect(ray1)
	var awnsr [2]float64

	awnsr[0] = 5
	awnsr[1] = 5

	if res[0].T != (awnsr[0]) || res[1].T != (awnsr[1]) {
		t.Fatalf("Tangent test failed")
	}

}

func TestIntersectionCenter(t *testing.T) {
	ray1 := ray.NewRay(*tuples.NewPoint(0, 0, 0), *tuples.NewVec3(1, 0, 0))

	sphere := NewSphere()

	var res Intersections = sphere.Intersect(ray1)
	var awnsr [2]float64

	awnsr[0] = -1
	awnsr[1] = 1

	if res[0].T != (awnsr[0]) || res[1].T != (awnsr[1]) {
		t.Fatalf("Tangent test failed")
	}

}

func TestIntersectionPast(t *testing.T) {
	ray1 := ray.NewRay(*tuples.NewPoint(5, 0, 0), *tuples.NewVec3(1, 0, 0))

	sphere := NewSphere()

	var res Intersections = sphere.Intersect(ray1)
	var awnsr [2]float64

	awnsr[0] = -6
	awnsr[1] = -4

	if res[0].T != (awnsr[0]) || res[1].T != (awnsr[1]) {
		t.Fatalf("Tangent test failed")
	}

}

func TestTransform(t *testing.T) {
	s := NewSphere()
	r := ray.NewRay(*tuples.NewPoint(0, 0, -5), *tuples.NewVec3(0, 0, 1))

	s.SetTransform(*matrix.Translate(0, 1, 2))

	res := s.Intersect(r)

	awns1 := float64(7)
	awns2 := float64(7)

	if res[0].T != (awns1) || res[1].T != awns2 {
		t.Fatalf("Fail")
	}
}

func TestLighting(t *testing.T) {
	sphere := NewSphere()
	position := tuples.NewPoint(0, 0, 0)
	sphere.Material = materials.DefaultMaterial()

	eyev := tuples.NewVec3(0, 1, -1)
	normalv := tuples.NewVec3(0, 0, -1)
	plight := light.WhitePointLight(*tuples.NewPoint(-10, 10, -10))

	l := light.Lighting(*position, *plight, sphere.Material, *eyev, *normalv, false)
	l.Print()

	if !l.Equals(view.NewColor(0.619615, 0.619615, 0.619615)) {
		t.Fatalf("Failed")
	}
}

func TestSphereRender(t *testing.T) {
	var canvas_pixels float64 = 150
	var wall_size float64 = 9

	ray_origin := tuples.NewPoint(0, 0, -5)

	pixel_size := float64(wall_size / canvas_pixels)
	half := wall_size / 2

	canvas := view.NewCanvas(int(canvas_pixels), int(canvas_pixels))
	//red := view.NewColor(1, 0, 0)
	black := view.NewColor(0, 0, 0)

	sphere := NewSphere()
	sphere.Material = materials.DefaultMaterial()
	sphere.Material.Color = *view.NewColor(1, 0, 1)
	//sphere.transform = *matrix.Scale(0.5, 1, 0.2)

	plight := light.WhitePointLight(*tuples.NewPoint(-10, 3, -10))
	for y := 0; y < int(canvas_pixels); y++ {

		world_y := half - pixel_size*float64(y)
		for x := 0; x < int(canvas_pixels); x++ {
			world_x := -half + pixel_size*float64(x)

			//calculate ray direction
			pos_screen := tuples.NewPoint(float64(world_x), float64(world_y), 10)
			sub := pos_screen.Sub(ray_origin)
			norm := sub.Normalize()
			//////
			r := ray.NewRay(*ray_origin, *norm)

			hits := sphere.Intersect(r)
			closest_hit := hits.Hit()

			if closest_hit.Object != nil {
				point := r.Position(closest_hit.T)
				normal := closest_hit.Object.NormalAt(point)
				eye := r.Direction.Negate()

				col := light.Lighting(*point, *plight, sphere.Material, *eye, *normal, false)
				canvas.WritePixel(x, y, *col)
			} else {
				canvas.WritePixel(x, y, *black)
			}
		}
	}

	err := canvas.WriteFile("Sphere")
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestReflection(t *testing.T) {
	s := NewSphere()
	s.Transform = *matrix.Translate(1, 1, 0)
	n := s.NormalAt(tuples.NewPoint(1, 2, 0))

	if !n.Equals(tuples.NewVec3(0, 1, 0)) {
		t.Fatalf("Incorect Noraml")
	}

	ray := tuples.NewVec3(1, -1, 0)
	r := ray.Reflect(*n)

	if !r.Equals(tuples.NewVec3(1, 1, 0)) {
		t.Fatalf("Incorect Reflect")
	}

}
