package main

type Config struct {
	ScreenWidth  int
	ScreenHeight int

	NearPlane float32
	FarPlane  float32

	MoveSpeed     float32
	RotationSpeed float64

	FOVStep    float64
	MinFOV     float64
	MaxFOV     float64
	DefaultFOV float64

	SceneFile string
}

func DefaultConfig() Config {
	return Config{
		ScreenWidth:   800,
		ScreenHeight:  800,
		NearPlane:     float32(0.1),
		FarPlane:      float32(100.0),
		MoveSpeed:     float32(0.05),
		RotationSpeed: float64(0.015),
		FOVStep:       float64(0.5),
		MinFOV:        float64(30.0),
		MaxFOV:        float64(120.0),
		DefaultFOV:    float64(90.0),
		SceneFile:     "blocks.json",
	}
}
