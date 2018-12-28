package charconv

import (
	"image"
	"image/color"

	"github.com/pbnjay/pixfont"
)

func ToBoolMap(msg string) [][]bool {
	height := 8
	textWidth := pixfont.MeasureString(msg)

	img := image.NewRGBA(image.Rect(0, 0, textWidth, height))
	pixfont.DrawString(img, 0, 0, msg, color.Black)

	boolMap := make([][]bool, height)
	for i := range boolMap {
		boolMap[i] = make([]bool, textWidth)
	}

	for y, row := range boolMap {
		for x, _ := range row {
			_, _, _, opacity  := img.At(x, y).RGBA()
			if opacity != 0 {
				boolMap[y][x] = true
			}
		}
	}

	return boolMap
}
