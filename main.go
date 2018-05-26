package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"github.com/octomarat/heroes-go/grid"
	"github.com/octomarat/heroes-go/unit"
	"github.com/octomarat/heroes-go/util"
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

	gr := grid.CreateGrid(gridPos(), gridWidth, gridHeight, cellSize)

	skeleton := unit.NewUnit(`resources/skeleton-small.png`, util.Position{X: 0, Y: 5}, gr)

	for !win.Closed() {
		win.Clear(colornames.Green)
		background.Draw(win, backMat)
		grid.DrawGrid(win, gr)
		skeleton.Draw(win)
		win.Update()
	}
}

func gridPos() pixel.Vec {
	return pixel.V((screenWidth-gridWidth*cellSize)/2, (screenHeight-gridHeight*cellSize)/2 * 0.6)
}

func main() {
	pixelgl.Run(run)
}
