package functions

import (
	"strings"
)

// ColorSubstring colors the specified substring within the concatenatedLines using the given colorCode.
func ColorSubstring(concatenatedLines []string, colorCode string, substring string, stringToColor string) []string {
	startColor := "\033[" + colorCode + "m"
	endColor := "\033[0m"
	coloredLines := make([]string, len(concatenatedLines))

	for i, line := range concatenatedLines {
		if substring == "" {
			coloredLines[i] = startColor + line + endColor
		} else {
			coloredLines[i] = strings.ReplaceAll(line, substring, startColor+substring+endColor)
		}
	}

	return coloredLines
}
