package object

import (
	"testing"
)

func TestHit(t *testing.T) {

	sphere := NewSphere()

	i1 := Intersection{T: 5, Object: sphere}
	i2 := Intersection{T: -2, Object: sphere}
	i3 := Intersection{T: 3, Object: sphere}
	i4 := Intersection{T: 1, Object: sphere}
	i5 := Intersection{T: 6, Object: sphere}

	inter := NewIntersections(i1, i2, i3, i4, i5)

	for _, intr := range inter {
		println(intr.T)
	}

	hit := inter.Hit()

	if hit != i4 {
		t.Fatalf("Failed")
	}
}

func TestHitNegativeT(t *testing.T) {

	sphere := NewSphere()

	i1 := Intersection{T: -1, Object: sphere}
	i2 := Intersection{T: 2, Object: sphere}

	inter := NewIntersections(i1, i2)

	hit := inter.Hit()

	if hit != i2 {
		t.Fatalf("Failed")
	}
}

func TestHitNegativeTAll(t *testing.T) {

	sphere := NewSphere()

	i1 := Intersection{T: -1, Object: sphere}
	i2 := Intersection{T: -2, Object: sphere}

	inter := NewIntersections(i1, i2)

	hit := inter.Hit()

	if hit == i2 || hit == i1 {
		t.Fatalf("Failed")
	}
}
