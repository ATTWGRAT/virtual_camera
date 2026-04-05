package main

type Vector4 struct {
	X, Y, Z, W float32
}

type Edge struct {
	Start int
	End   int
}

type Block struct {
	Vertices []Vector4
	Edges    []Edge
}

type Matrix4 [4][4]float32

func CreateProjectionMatrix(near, far float32) Matrix4 {
	var m Matrix4

	m[0][0] = 1.0
	m[1][1] = 1.0

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
