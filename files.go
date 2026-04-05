package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type BlockPlanesJSON struct {
	XMin float32 `json:"x_min"`
	XMax float32 `json:"x_max"`
	YMin float32 `json:"y_min"`
	YMax float32 `json:"y_max"`
	ZMin float32 `json:"z_min"`
	ZMax float32 `json:"z_max"`
}

type BlockJSON struct {
	ID     string          `json:"id"`
	Planes BlockPlanesJSON `json:"planes"`
}

type SceneJSON struct {
	Blocks []BlockJSON `json:"blocks"`
}

func loadSceneFromJSON(filename string) (*SceneJSON, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read JSON file: %w", err)
	}

	var scene SceneJSON

	err = json.Unmarshal(data, &scene)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return &scene, nil
}

func loadBlocksFromScene(scene *SceneJSON) []Block {
	var blocks []Block

	standardEdges := []Edge{
		{Start: 0, End: 1}, {Start: 1, End: 2}, {Start: 2, End: 3}, {Start: 3, End: 0}, // Near face
		{Start: 4, End: 5}, {Start: 5, End: 6}, {Start: 6, End: 7}, {Start: 7, End: 4}, // Far face
		{Start: 0, End: 4}, {Start: 1, End: 5}, {Start: 2, End: 6}, {Start: 3, End: 7}, // Connecting lines
	}

	for _, blockJSON := range scene.Blocks {
		planes := blockJSON.Planes

		vertices := []Vector4{
			{X: planes.XMin, Y: planes.YMin, Z: planes.ZMin, W: 1},
			{X: planes.XMax, Y: planes.YMin, Z: planes.ZMin, W: 1},
			{X: planes.XMax, Y: planes.YMax, Z: planes.ZMin, W: 1},
			{X: planes.XMin, Y: planes.YMax, Z: planes.ZMin, W: 1},
			{X: planes.XMin, Y: planes.YMin, Z: planes.ZMax, W: 1},
			{X: planes.XMax, Y: planes.YMin, Z: planes.ZMax, W: 1},
			{X: planes.XMax, Y: planes.YMax, Z: planes.ZMax, W: 1},
			{X: planes.XMin, Y: planes.YMax, Z: planes.ZMax, W: 1},
		}

		block := Block{
			Vertices: vertices,
			Edges:    standardEdges,
		}

		blocks = append(blocks, block)
	}

	return blocks
}

func LoadBlocks(filename string) ([]Block, error) {
	scene, err := loadSceneFromJSON(filename)
	if err != nil {
		return nil, err
	}

	blocks := loadBlocksFromScene(scene)
	return blocks, nil
}
