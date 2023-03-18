package main

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/TadaTeruki/NoiseGo/v2/noise"
)

func main() {

	nz := noise.New(30)

	width := 800
	height := 800
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	for iy := 0; iy < height; iy++ {
		for ix := 0; ix < width; ix++ {

			query_x := float64(ix) / float64(width) * 30
			query_y := float64(iy) / float64(height) * 30

			noise := nz.Get(query_x, query_y)

			color := color.RGBA{uint8(noise * 0xff), uint8(noise * 0xff), uint8(noise * 0xff), 0xff}
			img.Set(ix, iy, color)
		}
	}

	file, _ := os.Create("image.png")
	png.Encode(file, img)

}
