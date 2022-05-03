package view

import (
	"testing"
)

func TestCanvas(t *testing.T) {
	c1 := NewCanvas(10, 10)

	for i := 0; i < 10; i++ {
		for y := 0; y < 10; y++ {
			if !c1.data[i][y].Equals(NewColor(0, 0, 0)) {
				t.Fatalf("Error at index %d, %d", i, y)
			}
		}
	}
}

func TestCanvasColors(t *testing.T) {
	c1 := NewCanvas(10, 10)

	for i := 0; i < 10; i++ {
		for y := 0; y < 10; y++ {
			c1.WritePixel(i, y, *NewColor(0.2, 0, 0))
		}
	}

	col, _ := c1.PixelAt(1, 3)
	println(col.Red)

	if !c1.data[1][3].Equals(NewColor(0.2, 0, 0)) {
		t.Fatalf("Failed")
	}
}

func TestGenerateImage(t *testing.T) {
	c1 := NewCanvas(500, 500)

	c1.WritePixel(1, 0, *NewColor(0.2, 0, 0))
	c1.WritePixel(9, 0, *NewColor(0.2, 0.4, 0))
	c1.WritePixel(3, 4, *NewColor(0.2, 0, 1))

	c1.GenerateImage()
}

func TestSaveImage(t *testing.T) {
	c1 := NewCanvas(255, 255)

	for w := 0; w < c1.Width; w++ {
		for h := 0; h < c1.Height; h++ {
			c1.WritePixel(w, h, *NewColor(0, float64(h)/255, float64(w)/255))
		}
	}

	err := c1.WriteFile("Test")

	if err != nil {
		t.Fatalf(err.Error())
	}
}
