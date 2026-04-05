package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	Blocks []Block
}

func newGame(blocks []Block) *Game {
	return &Game{Blocks: blocks}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawCameraFrame(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
