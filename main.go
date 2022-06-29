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

	window1 := app1.NewWindow("Life!")

	// game := NewGame()
	// field := &CellField{Game: game}

	field := NewTriangleField()

	// pause button

	// pauseButton := widget.NewButton("Pause", func() {
	// 	game.TogglePause()
	// })
	// game.Paused.AddListener(binding.NewDataListener(func() {
	// 	p, _ := game.Paused.Get()
	// 	if p {
	// 		pauseButton.SetText("Unpause")
	// 	} else {
	// 		pauseButton.SetText("Pause")
	// 	}
	// }))

	// generation label

	// genBoundText := binding.IntToStringWithFormat(game.Generation, "Generation: %d")
	// generationText := widget.NewLabelWithData(genBoundText)
	bottomText := widget.NewLabel("sup")

	bottomToolbar := container.New(layout.NewGridLayout(2), bottomText)

	cont := container.New(
		layout.NewBorderLayout(nil, bottomToolbar, nil, nil),
		field,
		bottomToolbar,
	)

	go func() {
		for range time.Tick(time.Millisecond * 500) {
			field.Tick()
		}
	}()

	window1.SetContent(cont)
	window1.Resize(fyne.NewSize(600, 600))
	window1.SetFixedSize(true)
	window1.ShowAndRun()
}
