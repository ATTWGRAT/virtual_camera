package main

type RenderState struct {
	ViewMatrix       Matrix4
	ProjectionMatrix Matrix4

	lastCamera Camera
	lastFOV    float64
	ready      bool
}

func NewRenderState(camera Camera, fov float64, config Config) RenderState {
	var state RenderState
	state.Update(camera, fov, config)
	return state
}

func (r *RenderState) Update(camera Camera, fov float64, config Config) {
	if !r.ready || camera != r.lastCamera {
		r.ViewMatrix = CreateViewMatrix(camera)
		r.lastCamera = camera
	}

	if !r.ready || fov != r.lastFOV {
		r.ProjectionMatrix = CreateProjectionMatrix(config.NearPlane, config.FarPlane, fov)
		r.lastFOV = fov
	}

	r.ready = true
}

func CreateViewMatrix(camera Camera) Matrix4 {
	tMat := CreateTranslationMatrix(-camera.X, -camera.Y, -camera.Z)
	pMat := CreatePitchMatrix(-camera.Pitch)
	yMat := CreateYawMatrix(-camera.Yaw)
	rMat := CreateRollMatrix(-camera.Roll)

	return MultiplyMatrices(rMat, MultiplyMatrices(pMat, MultiplyMatrices(yMat, tMat)))
}
