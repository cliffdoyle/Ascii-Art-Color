package functions

import "fmt"

// PrintChar takes a rune and font data, and returns the corresponding lines for that character
func PrintChar(c rune, fontData [95][8]string) ([]string, error) {
    // Check if the character is within the supported ASCII range
    if c > 126 || c < 32 {
        return nil, fmt.Errorf("character out of range: %c", c)
    }

    // Calculate the index in the font data array
    charIndex := int(c - 32)

    // Preallocate the slice with a capacity of 8
    lines := make([]string, 8)

    // Populate the lines slice with the character's ASCII art
    for i := 0; i < 8; i++ {
        lines[i] = fontData[charIndex][i]
    }

    return lines, nil
}
