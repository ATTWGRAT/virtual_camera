package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Point2D struct {
	X, Y float32
}

func ClipLine(v1, v2 Vector4, nearPlane, farPlane float32) (Vector4, Vector4, bool) {
	// If both points are entirely behind the near plane or beyond the far plane, drop the line.
	if (v1.Z < nearPlane && v2.Z < nearPlane) || (v1.Z > farPlane && v2.Z > farPlane) {
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
	if v1.Z < nearPlane {
		v1 = interpolate(v1, v2, nearPlane)
	} else if v2.Z < nearPlane {
		v2 = interpolate(v1, v2, nearPlane)
	}

	// 2. Clip against the Far Plane
	if v1.Z > farPlane {
		v1 = interpolate(v1, v2, farPlane)
	} else if v2.Z > farPlane {
		v2 = interpolate(v1, v2, farPlane)
	}

	return v1, v2, true
}

func drawCameraFrame(blocks []Block, renderState RenderState, config Config, screen *ebiten.Image) {
	viewMatrix := renderState.ViewMatrix
	projection := renderState.ProjectionMatrix

	for _, block := range blocks {
		for _, edge := range block.Edges {
			v1 := block.Vertices[edge.Start]
			v2 := block.Vertices[edge.End]

			v1 = MultiplyMatrixVector(viewMatrix, v1)
			v2 = MultiplyMatrixVector(viewMatrix, v2)

			v1, v2, visible := ClipLine(v1, v2, config.NearPlane, config.FarPlane)
			if !visible {
				continue
			}

			p1 := MultiplyMatrixVector(projection, v1)
			p2 := MultiplyMatrixVector(projection, v2)

			ndcP1 := Point2D{X: p1.X / p1.W, Y: p1.Y / p1.W}
			ndcP2 := Point2D{X: p2.X / p2.W, Y: p2.Y / p2.W}

			screenX1 := (ndcP1.X + 1) * float32(config.ScreenWidth) / 2
			screenY1 := (1 - ndcP1.Y) * float32(config.ScreenHeight) / 2
			screenX2 := (ndcP2.X + 1) * float32(config.ScreenWidth) / 2
			screenY2 := (1 - ndcP2.Y) * float32(config.ScreenHeight) / 2

			vector.StrokeLine(screen, screenX1, screenY1, screenX2, screenY2, 2, color.White, false)
		}

	}
}
