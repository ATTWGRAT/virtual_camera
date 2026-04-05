package main

import "math"

type Vector4 struct {
	X, Y, Z, W float32
}

type Matrix4 [4][4]float32

func CreatePitchMatrix(a float64) Matrix4 {
	c, s := float32(math.Cos(a)), float32(math.Sin(a))
	var m Matrix4
	m[0][0] = 1
	m[1][1] = c
	m[1][2] = -s
	m[2][1] = s
	m[2][2] = c
	m[3][3] = 1
	return m
}

func CreateYawMatrix(a float64) Matrix4 {
	c, s := float32(math.Cos(a)), float32(math.Sin(a))
	var m Matrix4
	m[0][0] = c
	m[0][2] = s
	m[1][1] = 1
	m[2][0] = -s
	m[2][2] = c
	m[3][3] = 1
	return m
}

func CreateRollMatrix(a float64) Matrix4 {
	c, s := float32(math.Cos(a)), float32(math.Sin(a))
	var m Matrix4
	m[0][0] = c
	m[0][1] = -s
	m[1][0] = s
	m[1][1] = c
	m[2][2] = 1
	m[3][3] = 1
	return m
}

func CreateProjectionMatrix(near, far float32, fov float64) Matrix4 {
	var m Matrix4

	fovRadians := fov * (math.Pi / 180.0)

	f := 1.0 / float32(math.Tan(fovRadians/2.0))

	m[0][0] = f
	m[1][1] = f

	m[2][2] = far / (far - near)
	m[2][3] = (-far * near) / (far - near)

	m[3][2] = 1.0
	m[3][3] = 0.0

	return m
}

func MultiplyMatrixVector(m Matrix4, v Vector4) Vector4 {
	return Vector4{
		X: m[0][0]*v.X + m[0][1]*v.Y + m[0][2]*v.Z + m[0][3]*v.W,
		Y: m[1][0]*v.X + m[1][1]*v.Y + m[1][2]*v.Z + m[1][3]*v.W,
		Z: m[2][0]*v.X + m[2][1]*v.Y + m[2][2]*v.Z + m[2][3]*v.W,
		W: m[3][0]*v.X + m[3][1]*v.Y + m[3][2]*v.Z + m[3][3]*v.W,
	}
}

func MultiplyMatrices(A, B Matrix4) Matrix4 {
	var C Matrix4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C
}

// CreateTranslationMatrix creates a matrix to move points in 3D space
func CreateTranslationMatrix(tx, ty, tz float32) Matrix4 {
	var m Matrix4
	// Set the diagonal to 1.0 (Identity Matrix baseline)
	m[0][0] = 1.0
	m[1][1] = 1.0
	m[2][2] = 1.0
	m[3][3] = 1.0

	// Set the translation values in the far-right column
	m[0][3] = tx
	m[1][3] = ty
	m[2][3] = tz

	return m
}
