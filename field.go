package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

var (
	white    = color.White
	gridGray = color.Gray16{50000}
)

const (
	gridThicknessPixels int = 3
	cellSizePixels      int = 20
)

type Simulator interface {
	Tick()
	TogglePause()
}

type CellField struct {
	Game Simulator
	widget.BaseWidget

	dx int
	dy int

	renderer fyne.WidgetRenderer
}

func (g *CellField) CreateRenderer() fyne.WidgetRenderer {
	gameCanvas := canvas.NewRasterWithPixels(g.renderGrid)
	g.renderer = widget.NewSimpleRenderer(gameCanvas)
	return g.renderer
}

func (g *CellField) Tapped(e *fyne.PointEvent) {
	fmt.Println("yooo tapped")
	fmt.Printf("%+v\n", e)
}

func (g *CellField) Dragged(e *fyne.DragEvent) {
	// fmt.Println("dragged")
	// fmt.Printf("%+v\n", e)

	g.dx -= int(e.Dragged.DX)
	g.dy -= int(e.Dragged.DY)

	g.Refresh()
}

func (g *CellField) Refresh() {
	fmt.Println("in refresh...")
	g.BaseWidget.Refresh()
	g.renderer.Refresh()
}

func (g *CellField) DragEnd() {
	fmt.Println("we're done dragging")
}

func (g *CellField) renderGrid(x, y, _, _ int) color.Color {
	xx := x + g.dx
	yy := y + g.dy
	// fmt.Printf("renderGrid! x=%d xx=%d y=%d yy=%d\n", x, xx, y, yy)
	spacing := gridThicknessPixels + cellSizePixels
	if (0 < xx%spacing && xx%spacing <= gridThicknessPixels) || (0 < yy%spacing && yy%spacing <= gridThicknessPixels) {
		return gridGray
	}
	return white
}
