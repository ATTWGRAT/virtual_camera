package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func drawCameraFrame(screen *ebiten.Image) {
	cx, cy := float32(screenWidth)/2, float32(screenHeight)/2
	size := float32(100)
	strokeWidth := float32(2)
	strokeColor := color.White

	vector.StrokeLine(screen, cx-size, cy-size, cx+size, cy-size, strokeWidth, strokeColor, false)
	vector.StrokeLine(screen, cx+size, cy-size, cx+size, cy+size, strokeWidth, strokeColor, false)
	vector.StrokeLine(screen, cx+size, cy+size, cx-size, cy+size, strokeWidth, strokeColor, false)
	vector.StrokeLine(screen, cx-size, cy+size, cx-size, cy-size, strokeWidth, strokeColor, false)
}
