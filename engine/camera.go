package engine

import (
	"math"

	"github.com/gernivisser/raytracer/engine/ray"
	"github.com/gernivisser/raytracer/matrix"
	"github.com/gernivisser/raytracer/tuples"
	"github.com/gernivisser/raytracer/view"
)

type Camera struct {
	Hsize       int
	Vsize       int
	FOV         float64
	Transform   matrix.Matrix
	Pixel_size  float64
	half_width  float64
	half_heigth float64
}

func NewCamera(hsize int, vsize int, fov float64) *Camera {
	half_view := math.Tan(fov / 2)
	aspect := float64(hsize) / float64(vsize)

	var half_width, half_heigth float64

	if hsize >= vsize {
		half_width = half_view
		half_heigth = half_view / aspect
	} else {
		half_width = half_view * aspect
		half_heigth = half_view
	}

	return &Camera{
		Hsize:       hsize,
		Vsize:       vsize,
		Transform:   *matrix.Identity4,
		FOV:         fov,
		half_width:  half_width,
		half_heigth: half_heigth,
		Pixel_size:  (half_width * 2) / float64(hsize),
	}
}

func (c *Camera) ray_for_pixel(px, py int) ray.Ray {
	xoffset := (float64(px) + 0.5) * c.Pixel_size
	yoffset := (float64(py) + 0.5) * c.Pixel_size

	world_x := c.half_width - xoffset
	world_y := c.half_heigth - yoffset

	cinvtrans, succeeded := c.Transform.Invert()

	if !succeeded {
		panic("tranceform is not invertable")
	}

	pixel := *cinvtrans.MultiplyTuple(tuples.NewPoint(world_x, world_y, -1))
	origin := *cinvtrans.MultiplyTuple(tuples.NewPoint(0, 0, 0))
	direction := pixel.Sub(&origin).Normalize()

	return *ray.NewRay(origin, *direction)
}

func (camera *Camera) Render(world World) view.Canvas {
	image := view.NewCanvas(camera.Hsize, camera.Vsize)

	for y := 0; y <= camera.Vsize-1; y++ {
		for x := 0; x <= camera.Hsize-1; x++ {
			ray := camera.ray_for_pixel(x, y)
			color := world.Color_At(ray, 4)
			image.WritePixel(x, y, *color)
		}
	}
	return *image
}
