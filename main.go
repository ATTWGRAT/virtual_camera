package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	config := DefaultConfig()

	blocks, err := LoadBlocks(config.SceneFile)

	var camera Camera

	if err != nil {
		log.Fatalf("Error loading blocks: %v", err)
	}

	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)

	ebiten.SetWindowTitle("Virtual Camera")

	if err := ebiten.RunGame(newGame(blocks, camera, config)); err != nil {
		log.Fatal(err)
	}
}
