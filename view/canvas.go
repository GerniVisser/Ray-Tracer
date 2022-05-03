package view

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type Canvas struct {
	Width  int
	Height int
	data   [][]Color
}

func NewCanvas(width int, height int) *Canvas {
	data := make([][]Color, width)
	for i := 0; i < int(width); i++ {
		data[i] = make([]Color, height)
	}

	c := Canvas{Width: width, Height: height, data: data}
	return &c
}

func (c *Canvas) Print() {
	s := fmt.Sprintf("%d", c.Width) + " x " + fmt.Sprintf("%d", c.Height) + " Matrix\n"
	for x, row := range c.data {
		for col := range row {
			s += fmt.Sprintf("%0.2f", c.data[x][col]) + "\t| "
		}
		s += "\n"
	}
	fmt.Print(s)
}

func (c *Canvas) CheckPixelExist(x int, y int) bool {
	if c.Width < x || x < 0 {
		return false
	}
	if c.Height < y || y < 0 {
		return false
	}
	return true
}

func (c *Canvas) WritePixel(x int, y int, col Color) error {
	if !c.CheckPixelExist(x, y) {
		return errors.New("pixel coordinates out of bound for canvas")
	}

	c.data[x][y] = col
	return nil
}

func (c *Canvas) PixelAt(x int, y int) (*Color, error) {
	if !c.CheckPixelExist(x, y) {
		return nil, errors.New("pixel coordinates out of bound for canvas")
	}

	col := c.data[x][y]
	return &col, nil
}

func convertTo255(in float64) int {
	if in > 1 {
		in = 1
	}
	return int(in * 255)
}

func (c *Canvas) GenerateImage() string {

	var output string = "P3\n" + strconv.Itoa(c.Width) + " " + strconv.Itoa(c.Height) + "\n255\n"
	contents := make([]string, 0, c.Width*c.Height)

	for h := 0; h <= c.Height-1; h++ {
		for w := 0; w <= c.Width-1; w++ {
			contents = append(contents,
				strconv.Itoa(convertTo255(c.data[w][h].Red))+
					" "+strconv.Itoa(convertTo255(c.data[w][h].Green))+
					" "+strconv.Itoa(convertTo255(c.data[w][h].Blue))+"\n")
		}
	}

	output += strings.Join(contents, "")
	return output
}

func (c *Canvas) WriteFile(fileName string) error {
	path, _ := filepath.Abs("")

	for filepath.Base(path) != "Ray Tracer" {
		path = filepath.Dir(path)
	}
	path += "/temp/" + fileName + ".ppm"
	f, err := os.Create(path)

	if err != nil {
		return err
	}

	defer f.Close()

	image := c.GenerateImage()

	_, err2 := f.Write([]byte(image))

	if err2 != nil {
		return err2
	}

	//Open newly created file in GIMP
	cm := exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", path)
	err3 := cm.Start()

	if err3 != nil {
		return err3
	}

	return nil
}
