package main

import "github.com/hajimehoshi/ebiten/v2"

type InputCommand struct {
	Forward  float32
	Strafe   float32
	Vertical float32

	PitchDelta float64
	YawDelta   float64
	RollDelta  float64

	FOVDelta float64
}

type InputHandler struct {
	config Config
}

func NewInputHandler(config Config) InputHandler {
	return InputHandler{config: config}
}

func (h InputHandler) Read() InputCommand {
	var cmd InputCommand

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		cmd.Forward += h.config.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		cmd.Forward -= h.config.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		cmd.Strafe -= h.config.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		cmd.Strafe += h.config.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		cmd.Vertical += h.config.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		cmd.Vertical -= h.config.MoveSpeed
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		cmd.PitchDelta -= h.config.RotationSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		cmd.PitchDelta += h.config.RotationSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		cmd.YawDelta -= h.config.RotationSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		cmd.YawDelta += h.config.RotationSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyComma) {
		cmd.RollDelta += h.config.RotationSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyPeriod) {
		cmd.RollDelta -= h.config.RotationSpeed
	}

	if ebiten.IsKeyPressed(ebiten.KeyZ) {
		cmd.FOVDelta -= h.config.FOVStep
	}
	if ebiten.IsKeyPressed(ebiten.KeyX) {
		cmd.FOVDelta += h.config.FOVStep
	}

	return cmd
}

func ClampFOV(value float64, config Config) float64 {
	if value < config.MinFOV {
		return config.MinFOV
	}
	if value > config.MaxFOV {
		return config.MaxFOV
	}
	return value
}
