package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type grid struct {
	pos pixel.Vec
	size int
	cellSize float64
	cells [][]*cell
}

func createGrid(gridPos pixel.Vec, gridSize int, gridCellSize float64) *grid {
	return &grid{gridPos, gridSize, gridCellSize, createCells(gridPos, gridSize, gridCellSize)}
}

func createCells(gridPos pixel.Vec, gridSize int, gridCellSize float64) [][]*cell {
	cells := make([][]*cell, gridSize)
	for cx := 0; cx < gridSize; cx++ {
		cells[cx] = make([]*cell, gridSize)
		for cy := 0; cy < gridSize; cy++ {
			cellPos := gridPos.Add(pixel.V(float64(cx), float64(cy)).Scaled(gridCellSize))
			cells[cx][cy] = emptyCell(cellPos, gridCellSize)
		}
	}
	return cells
}

func drawGrid(window *pixelgl.Window, grid *grid) {
	var highlighted *cell
	mousePosition := window.MousePosition()
	for _, column := range grid.cells {
		for _, cell := range column {
			if cell.bounds.Contains(mousePosition) {
				highlighted = cell
			} else {
				cell.imdraw.Draw(window)
			}
		}
	}
	if highlighted != nil {
		highlightedCell(highlighted.bounds.Min, highlighted.bounds.W()).imdraw.Draw(window)
	}
}