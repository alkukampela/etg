package emojier

import (
	"strings"
)

const spacerEmoji = ":blank:"

func EmojifyBoolMap(boolMap [][]bool, emoji string) string {
	var builder strings.Builder
	for y, row := range boolMap {
		for x, _ := range row {
			if boolMap[y][x] {
				builder.WriteString(emoji)
			} else {
				builder.WriteString(spacerEmoji)
			}
		}
		builder.WriteString("\n")
	}

	return builder.String()
}
