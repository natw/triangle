package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// TriangleField is a field for a triangle
type TriangleField struct {
	Points    map[Coord]bool
	r         *rand.Rand
	Vertices  []Coord
	LastPoint Coord
	widget.BaseWidget
}

func NewTriangleField() *TriangleField {
	f := &TriangleField{}
	f.Vertices = []Coord{
		{300, 20},
		{20, 580},
		{580, 580},
	}
	f.Points = make(map[Coord]bool)
	f.LastPoint = Coord{300, 20}
	f.r = rand.New(rand.NewSource(time.Now().UnixNano()))

	return f
}

// CreateRenderer does something
func (f *TriangleField) CreateRenderer() fyne.WidgetRenderer {
	fmt.Println("what")
	canvas := canvas.NewRasterWithPixels(
		func(x, y, _, _ int) color.Color {
			toCheck := []Coord{
				{x, y},
				{x + 1, y},
				{x + 1, y - 1},
				{x + 1, y + 1},
				{x - 1, y},
				{x - 1, y - 1},
				{x - 1, y + 1},
				{x, y + 1},
				{x, y - 1},
			}
			for _, c := range toCheck {
				if f.Points[c] {
					return color.Black
				}
			}
			return color.White
		},
	)
	return widget.NewSimpleRenderer(canvas)
}

func (f *TriangleField) Tick() {
	i := f.r.Intn(3)
	target := f.Vertices[i]
	x := f.LastPoint.X + (target.X-f.LastPoint.X)/2
	y := f.LastPoint.Y + (target.Y-f.LastPoint.Y)/2
	c := Coord{x, y}
	f.Points[c] = true
	f.LastPoint = c
	fmt.Println(c)
}
