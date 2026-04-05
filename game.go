package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Blocks      []Block
	Camera      Camera
	FOV         float64
	Input       InputHandler
	Config      Config
	RenderState RenderState
}

func newGame(blocks []Block, camera Camera, config Config) *Game {
	fov := config.DefaultFOV
	return &Game{
		Blocks:      blocks,
		Camera:      camera,
		FOV:         fov,
		Input:       NewInputHandler(config),
		Config:      config,
		RenderState: NewRenderState(camera, fov, config),
	}
}

func (g *Game) Update() error {
	cmd := g.Input.Read()

	g.Camera.MoveForward(cmd.Forward)
	g.Camera.MoveRight(cmd.Strafe)
	g.Camera.MoveUp(cmd.Vertical)
	g.Camera.Rotate(cmd.PitchDelta, cmd.YawDelta, cmd.RollDelta)

	g.FOV = ClampFOV(g.FOV+cmd.FOVDelta, g.Config)
	g.RenderState.Update(g.Camera, g.FOV, g.Config)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawCameraFrame(g.Blocks, g.RenderState, g.Config, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Config.ScreenWidth, g.Config.ScreenHeight
}
