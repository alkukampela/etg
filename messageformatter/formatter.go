package messageformatter

import (
	"strings"
)

func Format(message string, responsiveMode bool) string {
	if responsiveMode {
		return separateCharsWithNewlines(message)
	}
	return message
}


func separateCharsWithNewlines(message string) string {
	var runes = []rune(message)
	var builder strings.Builder
	for _, rune := range runes {
		builder.WriteRune(rune)
		builder.WriteString("\n")
	}
	// Remove last newline from returned value
	return builder.String()[:builder.Len()-1]
}
