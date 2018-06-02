package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"github.com/octomarat/heroes-go/units"
	"github.com/octomarat/heroes-go/util"
	"github.com/octomarat/heroes-go/ground"
	"github.com/octomarat/heroes-go/game"
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

	gr := ground.CreateGrid(gridPos(), gridWidth, gridHeight, cellSize)

	skeleton := units.NewUnit(`resources/skeleton-35x42.png`, util.Position{X: 0, Y: 5}, false, gr)

	gameState := &game.State{}

	for !win.Closed() {
		win.Clear(colornames.Green)
		background.Draw(win, backMat)
		ground.DrawGrid(win, gr)

		var selectedPos *util.Position = nil
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			mousePosition := win.MousePosition()
			selectedPos = gr.CellPosition(mousePosition)
		}
		selected := false
		if gameState.IsSelected(skeleton) {
			selected = selectedPos == nil || skeleton.IsAtPosition(*selectedPos)
		} else {
			selected = selectedPos != nil && skeleton.IsAtPosition(*selectedPos)
		}
		skeleton.Draw(win, selected)
		if selected {
			gameState.SetSelectedUnit(skeleton)
		} else {
			gameState.SetSelectedUnit(nil)
		}

		win.Update()
	}
}

func gridPos() pixel.Vec {
	return pixel.V((screenWidth-gridWidth*cellSize)/2, (screenHeight-gridHeight*cellSize)/2 * 0.6)
}

func main() {
	pixelgl.Run(run)
}
