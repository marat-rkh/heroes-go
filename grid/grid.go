package grid

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/octomarat/heroes-go/cell"
)

type Grid struct {
	pos      pixel.Vec
	width    int
	height   int
	cellSize float64
	Cells    [][]*cell.Cell
}

func CreateGrid(pos pixel.Vec, width int, height int, cellSize float64) *Grid {
	return &Grid{
		pos:      pos,
		width:    width,
		height:   height,
		cellSize: cellSize,
		Cells:    createCells(pos, width, height, cellSize),
	}
}

func createCells(gridPos pixel.Vec, gridWidth int, gridHeight int, gridCellSize float64) [][]*cell.Cell {
	cells := make([][]*cell.Cell, gridWidth)
	for cx := 0; cx < gridWidth; cx++ {
		cells[cx] = make([]*cell.Cell, gridHeight)
		for cy := 0; cy < gridHeight; cy++ {
			cellPos := gridPos.Add(pixel.V(float64(cx), float64(cy)).Scaled(gridCellSize))
			cells[cx][cy] = cell.EmptyCell(cellPos, gridCellSize)
		}
	}
	return cells
}

func DrawGrid(window *pixelgl.Window, grid *Grid) {
	var highlighted *cell.Cell
	mousePosition := window.MousePosition()
	for _, column := range grid.Cells {
		for _, cell_ := range column {
			if cell_.Bounds.Contains(mousePosition) {
				highlighted = cell_
			} else {
				cell_.Imdraw.Draw(window)
			}
		}
	}
	if highlighted != nil {
		cell.HighlightedCell(highlighted.Bounds.Min, highlighted.Bounds.W()).Imdraw.Draw(window)
	}
}
