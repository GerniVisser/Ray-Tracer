package engine

import (
	"math"
	"testing"

	"github.com/gernivisser/raytracer/matrix"
	"github.com/gernivisser/raytracer/tuples"
	"github.com/gernivisser/raytracer/view"
)

func TestRaySize(t *testing.T) {
	c := NewCamera(200, 125, math.Pi/2)

	if c.Pixel_size != 0.01 {
		t.Fatalf("Error")
	}
}

func TestRayPixel(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	r := c.ray_for_pixel(0, 0)

	if !r.Origin.Equals(tuples.NewPoint(0, 0, 0)) || !r.Direction.Equals(tuples.NewVec3(0.66519, 0.33259, -0.66851)) {
		t.Fatalf("Error")
	}
}

func TestRayPixel2(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	c.Transform = *matrix.RotateY(math.Pi / 4).Multiply(matrix.Translate(0, -2, 5))
	r := c.ray_for_pixel(100, 50)

	if !r.Origin.Equals(tuples.NewPoint(0, 2, -5)) || !r.Direction.Equals(tuples.NewVec3(math.Sqrt(2)/2, 0, -math.Sqrt(2)/2)) {
		t.Fatalf("Error")
	}
}

func TestRayCameraWorld(t *testing.T) {
	w := DefaultWorld()
	c := NewCamera(11, 11, math.Pi/2)

	from := tuples.NewVec3(0, 0, -5)
	to := tuples.NewVec3(0, 0, 0)
	up := tuples.NewVec3(0, 1, 0)

	c.Transform = ViewTransform(*from, *to, *up)
	//c.Transform = *c.Transform.Multiply(matrix.Translate(0, 0, 3))
	image := c.Render(*w)

	col, _ := image.PixelAt(5, 5)
	image.WriteFile("test")

	if !col.Equals(view.NewColor(0.38066, 0.47583, 0.2855)) {
		col.Print()
		t.Fatalf("Error")
	}
}
