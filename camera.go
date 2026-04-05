package main

import "math"

// Camera holds the viewer position and orientation in world space.
type Camera struct {
	X, Y, Z          float32
	Pitch, Yaw, Roll float64
}

func (c *Camera) MoveForward(distance float32) {
	c.X += float32(math.Sin(c.Yaw)) * distance
	c.Z += float32(math.Cos(c.Yaw)) * distance
}

func (c *Camera) MoveRight(distance float32) {
	c.X += float32(math.Cos(c.Yaw)) * distance
	c.Z += float32(-math.Sin(c.Yaw)) * distance
}

func (c *Camera) MoveUp(distance float32) {
	c.Y += distance
}

func (c *Camera) Rotate(pitchDelta, yawDelta, rollDelta float64) {
	c.Pitch += pitchDelta
	c.Yaw += yawDelta
	c.Roll += rollDelta
}
