package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app1 := app.New()

	window1 := app1.NewWindow("a triangle?")

	field := NewTriangleField()

	bottomText := widget.NewLabel("sup")

	bottomToolbar := container.New(layout.NewGridLayout(2), bottomText)

	cont := container.New(
		layout.NewBorderLayout(nil, bottomToolbar, nil, nil),
		field,
		bottomToolbar,
	)

	go func() {
		time.Sleep(time.Second)
		for range time.Tick(time.Millisecond * 300) {
			field.Tick()
			field.Renderer.Refresh()
		}
	}()

	window1.SetContent(cont)
	window1.Resize(fyne.NewSize(600, 600))
	window1.SetFixedSize(true)
	window1.ShowAndRun()
}

type Coord struct {
	X, Y int
}
