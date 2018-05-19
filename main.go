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

	grid := createGrid(pixel.V(150, 50), 10, 50)
	for !win.Closed() {
		win.Clear(colornames.Aliceblue)
		drawGrid(win, grid)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
