package charconv

import (
	"image"
	"image/color"
	"strings"

	"github.com/pbnjay/pixfont"
)

const CharHeight = 8

func MessageToBoolMap(message string) [][]bool {
	image := createCanvas(message)
	drawMessageToImage(message, image)
	return imageToBoolMap(image)
}

func createCanvas(message string) *image.RGBA {
	height := countRows(message) * CharHeight
	width := longestWordWidth(message)
	return image.NewRGBA(image.Rect(0, 0, width, height))
}

func countRows(message string) int {
	return strings.Count(message, "\n") + 1
}

func longestWordWidth(message string) int {
	words := strings.Split(message, "\n")
	longestWord := 0
	for _, word := range words {
		if (pixfont.MeasureString(word) > longestWord) {
			longestWord = pixfont.MeasureString(word)
		}
	}
	return longestWord
}

func drawMessageToImage(message string, targetImage pixfont.Drawable) {
	words := strings.Split(message, "\n")
	for row, word := range words {
		pixfont.DrawString(targetImage, 0, row*CharHeight, word, color.Black)
	}
}

func imageToBoolMap(image *image.RGBA) [][]bool {
	boolMap := make([][]bool, image.Bounds().Max.Y)
	for i := range boolMap {
		boolMap[i] = make([]bool, image.Bounds().Max.X)
	}

	for y, row := range boolMap {
		for x, _ := range row {
			_, _, _, opacity  := image.At(x, y).RGBA()
			if opacity != 0 {
				boolMap[y][x] = true
			}
		}
	}

	return boolMap
}
