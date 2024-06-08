package functions

import (
	"bufio"
	"os"
)

// Fonts loads font data from a specified text file
// The function expects the file name without an extension
func Fonts(s string) ([95][8]string, error) {
	s += ".txt"

	var fontData [95][8]string

	file, err := os.Open(s)
	if err != nil {
		return [95][8]string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	currentCha := -1
	currentLine := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// An empty line indicates the end of a character block
			currentCha++
			currentLine = 0
		} else {
			// Fill the current character's line
			fontData[currentCha][currentLine] = line
			currentLine++
		}
	}

	if scanner.Err() != nil {
		return [95][8]string{}, scanner.Err()
	}

	return fontData, nil
}
