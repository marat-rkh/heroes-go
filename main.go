package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
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
	background.Draw(win, backMat)
	grid := createGrid(pixel.V(150, 50), 10, 50)
	for !win.Closed() {
		drawGrid(win, grid)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
