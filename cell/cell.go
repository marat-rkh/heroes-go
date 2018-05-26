package cell

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
	"image/color"
)

type Cell struct {
	Bounds pixel.Rect
	Imdraw *imdraw.IMDraw
}

func EmptyCell(pos pixel.Vec, size float64) *Cell {
	return &Cell{
		Bounds: pixel.R(pos.X, pos.Y, pos.X+size, pos.Y+size),
		Imdraw: cellImdraw(pos, size, colornames.Darkgreen, 1, pixel.RGB(0, 255, 0).Mul(pixel.Alpha(0.005))),
	}
}

func HighlightedCell(pos pixel.Vec, size float64) *Cell {
	return &Cell{
		Bounds: pixel.R(pos.X, pos.Y, pos.X+size, pos.Y+size),
		Imdraw: cellImdraw(pos, size, colornames.Darkgreen, 2, pixel.RGB(0, 255, 0).Mul(pixel.Alpha(0.5))),
	}
}

func cellImdraw(
	pos pixel.Vec,
	size float64,
	borderColor color.Color,
	borderThickness float64,
	fillColor color.Color,
) *imdraw.IMDraw {
	cell := imdraw.New(nil)

	cell.Color = fillColor
	cell.Push(pos, pos.Add(pixel.V(size, size)))
	cell.Rectangle(0)

	cell.Color = borderColor
	cell.Push(pos, pos.Add(pixel.V(size, size)))
	cell.Rectangle(borderThickness)
	return cell
}
