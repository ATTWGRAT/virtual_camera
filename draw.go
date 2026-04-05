package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Point2D struct {
	X, Y float32
}

func ClipLine(v1, v2 Vector4) (Vector4, Vector4, bool) {
	// If both points are entirely behind the near plane or beyond the far plane, drop the line.
	if (v1.Z < near && v2.Z < near) || (v1.Z > far && v2.Z > far) {
		return v1, v2, false
	}

	// Helper function to interpolate a point at a specific target Z
	interpolate := func(p1, p2 Vector4, targetZ float32) Vector4 {
		t := (targetZ - p1.Z) / (p2.Z - p1.Z)
		return Vector4{
			X: p1.X + t*(p2.X-p1.X),
			Y: p1.Y + t*(p2.Y-p1.Y),
			Z: targetZ,
			W: 1.0,
		}
	}

	// 1. Clip against the Near Plane
	if v1.Z < near {
		v1 = interpolate(v1, v2, near)
	} else if v2.Z < near {
		v2 = interpolate(v1, v2, near)
	}

	// 2. Clip against the Far Plane
	if v1.Z > far {
		v1 = interpolate(v1, v2, far)
	} else if v2.Z > far {
		v2 = interpolate(v1, v2, far)
	}

	return v1, v2, true
}

func drawCameraFrame(g *Game, screen *ebiten.Image) {
	tMat := CreateTranslationMatrix(-g.Camera.X, -g.Camera.Y, -g.Camera.Z)
	pMat := CreatePitchMatrix(-g.Camera.Pitch)
	yMat := CreateYawMatrix(-g.Camera.Yaw)
	rMat := CreateRollMatrix(-g.Camera.Roll)

	// 2. Combine them (Order: Roll * Pitch * Yaw * Translation)
	// This order ensures we rotate around the camera's local origin
	viewMatrix := MultiplyMatrices(rMat, MultiplyMatrices(pMat, MultiplyMatrices(yMat, tMat)))

	for _, block := range g.Blocks {
		for _, edge := range block.Edges {
			v1 := block.Vertices[edge.Start]
			v2 := block.Vertices[edge.End]

			v1 = MultiplyMatrixVector(viewMatrix, v1)
			v2 = MultiplyMatrixVector(viewMatrix, v2)

			v1, v2, visible := ClipLine(v1, v2)
			if !visible {
				continue
			}

			p1 := MultiplyMatrixVector(g.ProjectionMatrix, v1)
			p2 := MultiplyMatrixVector(g.ProjectionMatrix, v2)

			ndcP1 := Point2D{X: p1.X / p1.W, Y: p1.Y / p1.W}
			ndcP2 := Point2D{X: p2.X / p2.W, Y: p2.Y / p2.W}

			screenX1 := (ndcP1.X + 1) * float32(screenWidth) / 2
			screenY1 := (1 - ndcP1.Y) * float32(screenHeight) / 2
			screenX2 := (ndcP2.X + 1) * float32(screenWidth) / 2
			screenY2 := (1 - ndcP2.Y) * float32(screenHeight) / 2

			vector.StrokeLine(screen, screenX1, screenY1, screenX2, screenY2, 2, color.White, false)
		}

	}
}
