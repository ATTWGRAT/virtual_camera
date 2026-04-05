package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Point2D struct {
	X, Y float32
}

func drawCameraFrame(g *Game, screen *ebiten.Image) {
	for _, block := range g.Blocks {

		projectedVertices := make([]Point2D, len(block.Vertices))

		for i, v := range block.Vertices {
			projected := MultiplyMatrixVector(g.ProjectionMatrix, v)

			if projected.W <= 0 {
				projected.W = 0.0001
			}

			ndcX := projected.X / projected.W
			ndcY := projected.Y / projected.W

			screenX := (ndcX + 1) * screenWidth / 2
			screenY := (1 - ndcY) * screenHeight / 2

			projectedVertices[i] = Point2D{X: screenX, Y: screenY}

		}

		for _, edge := range block.Edges {
			start := projectedVertices[edge.Start]
			end := projectedVertices[edge.End]

			vector.StrokeLine(screen, start.X, start.Y, end.X, end.Y, 2, color.White, false)
		}
	}
}
