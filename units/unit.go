package units

import (
	"github.com/faiface/pixel"
	"github.com/octomarat/heroes-go/util"
	"github.com/faiface/pixel/pixelgl"
	"github.com/octomarat/heroes-go/ground"
)

type Unit struct {
	spritePath string
	sprite     *pixel.Sprite
	cell       *ground.Cell
	gridPos    util.Position
	isLarge    bool
	screenPos  pixel.Vec
}

func NewUnit(spritePath string, gridPos util.Position, isLarge bool, grid *ground.Grid) *Unit {
	spritePic, err := util.LoadPicture(spritePath)
	if err != nil {
		panic(err)
	}
	gridCell := grid.Cells[gridPos.X][gridPos.Y]
	unitCellSize := gridCell.Bounds.H()
	if isLarge {
		unitCellSize *= 2
	}
	screenPos := gridCell.Bounds.Min.Add(spritePic.Bounds().Size().Scaled(0.5))
	return &Unit{
		spritePath: spritePath,
		sprite:     pixel.NewSprite(spritePic, spritePic.Bounds()),
		cell:       ground.SelectedCell(gridCell.Bounds.Min, unitCellSize),
		gridPos:    gridPos,
		isLarge:    isLarge,
		screenPos:  screenPos,
	}
}

func (u *Unit) Draw(window *pixelgl.Window, selected bool) {
	if selected {
		u.cell.Imdraw.Draw(window)
	}
	u.sprite.Draw(window, pixel.IM.Moved(u.screenPos))
}

func (u *Unit) IsAtPosition(position util.Position) bool {
	if position == u.gridPos {
		return true
	}
	if u.isLarge {

	}
	return false
}