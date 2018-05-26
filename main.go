package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:     "HoMM Go",
		Bounds:    pixel.R(0, 0, 800, 600),
		VSync:     true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	background := backgroundSprite()
	backMat := pixel.IM
	backMat = backMat.Scaled(pixel.ZV, 0.4)
	backMat = backMat.Moved(win.Bounds().Center().Add(pixel.V(0, 40)))
	grid := createGrid(pixel.V(50, 50), 14, 7, 50)
	for !win.Closed() {
		win.Clear(colornames.Green)
		background.Draw(win, backMat)
		drawGrid(win, grid)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
