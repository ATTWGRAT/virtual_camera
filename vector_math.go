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
