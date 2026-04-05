package main

type Edge struct {
	Start int
	End   int
}

type Block struct {
	Vertices []Vector4
	Edges    []Edge
}
