package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
	"image/color"
)

type cell struct {
	bounds pixel.Rect
	imdraw *imdraw.IMDraw
}

func emptyCell(pos pixel.Vec, size float64) *cell {
	return &cell{
		bounds: pixel.R(pos.X, pos.Y, pos.X+size, pos.Y+size),
		imdraw: cellImdraw(pos, size, colornames.Darkgreen, 1, pixel.RGB(0, 255, 0).Mul(pixel.Alpha(0.005))),
	}
}

func highlightedCell(pos pixel.Vec, size float64) *cell {
	return &cell{
		bounds: pixel.R(pos.X, pos.Y, pos.X+size, pos.Y+size),
		imdraw: cellImdraw(pos, size, colornames.Darkgreen, 2, pixel.RGB(0, 255, 0).Mul(pixel.Alpha(0.5))),
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
