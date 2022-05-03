package ray

import (
	"testing"

	"github.com/gernivisser/raytracer/matrix"
	"github.com/gernivisser/raytracer/tuples"
)

func TestEqual(t *testing.T) {
	a := NewRay(*tuples.NewPoint(0, 0, 0), *tuples.NewVec3(1, 0, 0))

	res := a.Position(0.5)

	awns := tuples.NewPoint(0.5, 0, 0)

	if !res.Equals(awns) {
		t.Fatalf("Fail")
	}

}

func TestRayTransform(t *testing.T) {
	r := NewRay(*tuples.NewPoint(1, 2, 3), *tuples.NewVec3(0, 1, 0))

	m := matrix.Translate(3, 4, 5)
	r1 := r.Transform(m)

	awns1 := tuples.NewPoint(4, 6, 8)
	awns2 := tuples.NewVec3(0, 1, 0)

	println(r.Origin.X)

	if r1.Origin.Equals(&r.Origin) {
		t.Fatalf("Fail")
	}

	if !r1.Origin.Equals(awns1) || !r1.Direction.Equals(awns2) {
		t.Fatalf("Fail")
	}
}

func TestRayTransformScale(t *testing.T) {
	r := NewRay(*tuples.NewPoint(1, 2, 3), *tuples.NewVec3(0, 1, 0))

	m := matrix.Scale(2, 3, 4)
	r1 := r.Transform(m)

	awns1 := tuples.NewPoint(2, 6, 12)
	awns2 := tuples.NewVec3(0, 3, 0)

	if r1.Origin.Equals(&r.Origin) || r1.Direction.Equals(&r.Direction) {
		t.Fatalf("Fail")
	}

	if !r1.Origin.Equals(awns1) || !r1.Direction.Equals(awns2) {
		t.Fatalf("Fail")
	}

}
