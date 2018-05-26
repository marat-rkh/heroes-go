package unit

import (
	"github.com/faiface/pixel"
	"github.com/octomarat/heroes-go/util"
	"github.com/faiface/pixel/pixelgl"
	"github.com/octomarat/heroes-go/grid"
)

type Unit struct {
	spritePath string
	sprite     *pixel.Sprite
	GridPos    util.Position
	ScreenPos  pixel.Vec
}

func NewUnit(spritePath string, gridPos util.Position, grid *grid.Grid) *Unit {
	spritePic, err := util.LoadPicture(spritePath)
	if err != nil {
		panic(err)
	}
	return &Unit{
		spritePath: spritePath,
		sprite:     pixel.NewSprite(spritePic, spritePic.Bounds()),
		GridPos:    gridPos,
		ScreenPos:  grid.Cells[gridPos.X][gridPos.Y].Bounds.Min.Add(spritePic.Bounds().Size().Scaled(0.5)),
	}
}

func (u *Unit) Draw(window *pixelgl.Window) {
	u.sprite.Draw(window, pixel.IM.Moved(u.ScreenPos))
}
