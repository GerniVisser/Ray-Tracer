package object

import (
	"testing"

	"github.com/gernivisser/raytracer/matrix"
)

func TestObjectSphere(t *testing.T) {
	s := NewSphere()

	if s.Transform.Equals(matrix.Identity4) {
		t.Fatalf("Error transform")
	}
}
