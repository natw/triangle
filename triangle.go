package main

import (
	"image/color"
	"math/rand"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// TriangleField is a field for a triangle
type TriangleField struct {
	Renderer   fyne.WidgetRenderer
	Points     map[Coord]bool
	pointMutex *sync.RWMutex
	r          *rand.Rand
	Vertices   []Coord
	LastPoint  Coord
	widget.BaseWidget
}

func (f *TriangleField) AddPoint(p Coord) {
	f.pointMutex.Lock()
	defer f.pointMutex.Unlock()
	f.Points[p] = true
}

func (f *TriangleField) HasPoint(p Coord) bool {
	// fmt.Println("in HasPoint")
	// fmt.Println(f.pointMutex)
	f.pointMutex.RLock()
	defer f.pointMutex.RUnlock()
	return f.Points[p]
}

func NewTriangleField() *TriangleField {
	f := &TriangleField{
		Vertices: []Coord{
			{600, 20},
			{20, 1180},
			{1180, 1180},
		},
		Points:     make(map[Coord]bool),
		pointMutex: &sync.RWMutex{},
		LastPoint:  Coord{600, 20},
		r:          rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	return f
}

// CreateRenderer does something
func (f *TriangleField) CreateRenderer() fyne.WidgetRenderer {
	canvas := canvas.NewRasterWithPixels(
		func(x, y, _, _ int) color.Color {
			// c := Coord{x, y}
			// if f.Points[c] {
			// 	return color.Black
			// }
			// return color.White

			vcs := []Coord{
				{599, 19},
				{599, 20},
				{599, 21},
				{600, 19},
				{600, 20},
				{600, 21},
				{601, 19},
				{601, 20},
				{601, 21},

				{19, 1179},
				{19, 1180},
				{19, 1181},
				{20, 1179},
				{20, 1180},
				{20, 1181},
				{21, 1179},
				{21, 1180},
				{21, 1181},

				{1179, 1179},
				{1180, 1180},
				{1181, 1181},
				{1179, 1179},
				{1180, 1180},
				{1181, 1181},
				{1179, 1179},
				{1180, 1180},
				{1181, 1181},
			}
			for _, v := range vcs {
				if v.X == x && v.Y == y {
					return color.RGBA{0xff, 0x00, 0x00, 0xff}
				}
			}

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
				if f.HasPoint(c) {
					return color.Black
				}
			}
			return color.White
		},
	)
	f.Renderer = widget.NewSimpleRenderer(canvas)
	return f.Renderer
}

func (f *TriangleField) Refresh() {
	f.Renderer.Refresh()
}

func (f *TriangleField) Tick() {
	i := f.r.Intn(3)
	target := f.Vertices[i]
	x := f.LastPoint.X + (target.X-f.LastPoint.X)/2
	y := f.LastPoint.Y + (target.Y-f.LastPoint.Y)/2
	c := Coord{x, y}
	f.AddPoint(c)
	f.LastPoint = c
}
