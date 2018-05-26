package main

import (
	"github.com/faiface/pixel"
	_ "image/jpeg"
	"github.com/octomarat/heroes-go/util"
)

func backgroundSprite() *pixel.Sprite {
	pic, err := util.LoadPicture("resources/field0.jpg")
	if err != nil {
		panic(err)
	}
	return pixel.NewSprite(pic, pic.Bounds())
}
