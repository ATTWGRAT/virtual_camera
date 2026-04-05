package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	blocks, err := LoadBlocks("blocks.json")

	projectionMatrix := CreateProjectionMatrix(near, far)

	if err != nil {
		log.Fatalf("Error loading blocks: %v", err)
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)

	ebiten.SetWindowTitle("Virtual Camera")

	if err := ebiten.RunGame(newGame(blocks, projectionMatrix)); err != nil {
		log.Fatal(err)
	}
}
