package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type grid struct {
	pos      pixel.Vec
	width    int
	height   int
	cellSize float64
	cells    [][]*cell
}

func createGrid(pos pixel.Vec, width int, height int, cellSize float64) *grid {
	return &grid{
		pos:      pos,
		width:    width,
		height:   height,
		cellSize: cellSize,
		cells:    createCells(pos, width, height, cellSize),
	}
}

func createCells(gridPos pixel.Vec, gridWidth int, gridHeight int, gridCellSize float64) [][]*cell {
	cells := make([][]*cell, gridWidth)
	for cx := 0; cx < gridWidth; cx++ {
		cells[cx] = make([]*cell, gridHeight)
		for cy := 0; cy < gridHeight; cy++ {
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
