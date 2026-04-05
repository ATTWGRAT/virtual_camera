package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	Blocks           []Block
	ProjectionMatrix Matrix4
}

func newGame(blocks []Block, projectionMatrix Matrix4) *Game {
	return &Game{Blocks: blocks, ProjectionMatrix: projectionMatrix}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawCameraFrame(g, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
