package main

import (
	"fyne.io/fyne/v2/data/binding"
)

func NewGame() *Game {
	g := &Game{
		Paused:     binding.NewBool(),
		Generation: binding.NewInt(),
	}

	g.cells = make(map[Coord]bool)
	g.cells[Coord{1, 1}] = true
	g.cells[Coord{3, 1}] = true
	return g
}

type Coord struct {
	X, Y int
}

type Game struct {
	Paused     binding.Bool
	Generation binding.Int
	cells      map[Coord]bool
}

func (g *Game) Tick() {
}

func (g *Game) TogglePause() {
	p, _ := g.Paused.Get()
	_ = g.Paused.Set(!p)
}
