package main

import (
	"github.com/faiface/pixel"
	"os"
	"image"
	_ "image/jpeg"
)

func backgroundSprite() *pixel.Sprite {
	pic, err := loadPicture("resources/field0.jpg")
	if err != nil {
		panic(err)
	}
	return pixel.NewSprite(pic, pic.Bounds())
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}