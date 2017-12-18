package render

import (
	"image"
	"image/color"
	"image/png"
	"io"
)

func ToPNG(target io.Writer) error {
	return png.Encode(target, createNeutralGrayPng())
}

func createNeutralGrayPng() image.Image {
	const width, height = 256, 256

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(128),
				G: uint8(128),
				B: uint8(128),
				A: 255,
			})
		}
	}

	return img
}
