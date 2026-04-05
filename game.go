package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Blocks           []Block
	ProjectionMatrix Matrix4
	Camera           Camera
	FOV              float64
}

func newGame(blocks []Block, camera Camera) *Game {
	return &Game{Blocks: blocks, ProjectionMatrix: CreateProjectionMatrix(screenWidth, screenHeight, defaultFOV), Camera: camera, FOV: defaultFOV}
}

func (g *Game) Update() error {
	// 1. Calculate the direction vectors based on Yaw
	// Math.Sin/Cos expect Radians.
	// Forward Vector:
	forwardX := float32(math.Sin(g.Camera.Yaw))
	forwardZ := float32(math.Cos(g.Camera.Yaw))

	// Right Vector (always 90 degrees perpendicular to Forward):
	rightX := float32(math.Cos(g.Camera.Yaw))
	rightZ := float32(-math.Sin(g.Camera.Yaw))

	// 2. Apply Movement relative to these vectors
	// Forward / Backward (W/S)
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Camera.X += forwardX * speed
		g.Camera.Z += forwardZ * speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Camera.X -= forwardX * speed
		g.Camera.Z -= forwardZ * speed
	}

	// Left / Right Strafe (A/D)
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Camera.X -= rightX * speed
		g.Camera.Z -= rightZ * speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Camera.X += rightX * speed
		g.Camera.Z += rightZ * speed
	}

	// Up / Down (Y-axis)
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.Camera.Y += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		g.Camera.Y -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Camera.Pitch -= rotSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Camera.Pitch += rotSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.Camera.Yaw -= rotSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.Camera.Yaw += rotSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyComma) {
		g.Camera.Roll += rotSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyPeriod) {
		g.Camera.Roll -= rotSpeed
	}

	if ebiten.IsKeyPressed(ebiten.KeyZ) { // Zoom In
		g.FOV -= 0.5
	}
	if ebiten.IsKeyPressed(ebiten.KeyX) { // Zoom Out
		g.FOV += 0.5
	}

	if g.FOV < 30 {
		g.FOV = 30
	}
	if g.FOV > 120 {
		g.FOV = 120
	}

	g.ProjectionMatrix = CreateProjectionMatrix(800, 800, g.FOV)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawCameraFrame(g, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
