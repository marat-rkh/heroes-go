package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

const (
	gridWidth  = 20
	gridHeight = 8
	cellSize   = screenWidth / gridWidth * 0.8
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "HoMM Go",
		Bounds: pixel.R(0, 0, screenWidth, screenHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	background := backgroundSprite()
	backMat := pixel.IM
	backMat = backMat.Scaled(pixel.ZV, 0.4)
	backMat = backMat.Moved(win.Bounds().Center().Add(pixel.V(0, 40)))

	grid := createGrid(gridPos(), gridWidth, gridHeight, cellSize)
	for !win.Closed() {
		win.Clear(colornames.Green)
		background.Draw(win, backMat)
		drawGrid(win, grid)
		win.Update()
	}
}

func gridPos() pixel.Vec {
	return pixel.V((screenWidth-gridWidth*cellSize)/2, (screenHeight-gridHeight*cellSize)/2 * 0.6)
}

func main() {
	pixelgl.Run(run)
}
