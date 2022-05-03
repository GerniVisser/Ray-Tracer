package object

import (
	"math"
	"testing"

	"github.com/gernivisser/raytracer/engine/ray"
	"github.com/gernivisser/raytracer/matrix"
	"github.com/gernivisser/raytracer/tuples"
)

func TestNormal(t *testing.T) {
	p := NewPlane()
	p.Transform = *matrix.RotateZ(math.Pi / 2)

	norm := p.NormalAt(tuples.NewPoint(-1, 0, 1))

	if !norm.Equals(tuples.NewPoint(-1, 0, 0)) {
		t.Fatalf("Middle test failed")
	}
}

func TestIntersection(t *testing.T) {
	p := NewPlane()

	r := ray.NewRay(*tuples.NewPoint(0, 0, -5), *tuples.NewVec3(0, 0, 0))
	p.Transform = *matrix.RotateZ(math.Pi / 2)

	intersections := p.Intersect(r)

	if len(intersections) != 0 {
		t.Fatalf("Parallel test failed")
	}

	r = ray.NewRay(*tuples.NewPoint(1, 0, 0), *tuples.NewVec3(-1, 0, 0))

	intersections = p.Intersect(r)

	if intersections.Hit().T != 1 {
		t.Fatalf("Above test failed")
	}

	r = ray.NewRay(*tuples.NewPoint(1, 0, 0), *tuples.NewVec3(-1, 0, 0))

	intersections = p.Intersect(r)

	if intersections.Hit().T != 1 {
		t.Fatalf("Below test failed")
	}
}
