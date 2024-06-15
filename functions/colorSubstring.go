package functions

import (
	"fmt"
	"log"
)

// ColorSubstring colors the specified substring within the ASCII art lines.
func ColorSubstring(sub string, fontdata [95][8]string, color string) []string {
	concatenatedLines := make([]string, 8)

	for _, char := range sub {
		asciiArt, err := PrintChar(char, fontdata)
		if err != nil {
			log.Fatal(err)
		}
		for k, art := range asciiArt {

			if color != "" {
				concatenatedLines[k] += fmt.Sprintf("\033[%sm%s\033[0m", color, art)
			} else {
				concatenatedLines[k] += art
			}

		}
	}

	return concatenatedLines
}
