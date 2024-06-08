package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"ascii-art/functions"
)

// Map of color names to ANSI escape codes
var colorMap = map[string]string{
	"black":   "30",
	"red":     "31",
	"green":   "32",
	"yellow":  "33",
	"blue":    "34",
	"magenta": "35",
	"cyan":    "36",
	"white":   "37",
}

func main() {
	// Define flags
	outputFlag := flag.String("output", "", "Specify file type")
	colorFlag := flag.String("color", "", "Specify color and substring to be colored in the format <color> <substring>")
	flag.Parse()

	// Get non-flag arguments
	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("Usage: go run . [OPTION] [STRING] [FONT_FILE]")
		return
	}

	inputString := args[0] // Corrected index to 0 to get the string to be colored
	var fontFile string
	if len(args) == 3 { // Adjust for font file argument, if provided
		fontFile = args[2]
	} else {
		fontFile = "standard"
	}

	// Load font data
	fontArray, err := functions.Fonts(fontFile)
	if err != nil {
		fmt.Println("Error loading font:", err)
		return
	}

	// Prepare input string
	inputString = strings.ReplaceAll(inputString, "\\n", "\n")
	words := strings.Split(inputString, "\n")

	// Determine output destination
	var output *os.File
	if *outputFlag != "" {
		output, err = os.Create(*outputFlag)
		if err != nil {
			log.Fatal("Error creating output file:", err)
		}
		defer output.Close()
	} else {
		output = os.Stdout
	}

	// Extract substring and string to be colored
	var substring string
	var stringToColor string
	if *colorFlag != "" {
		colorParts := strings.SplitN(*colorFlag, " ", 2)
		if len(colorParts) != 2 {
			log.Fatal("Usage: go run . [OPTION] [STRING]")
			return
		}
		substring = colorParts[1] // Corrected index to 1 to get the substring
		stringToColor = args[0]   // Corrected index to 0 to get the string to be colored
	}

	for _, word := range words {
		concatenatedLines := make([]string, 8)
		count := 0
		if word == "" {
			count++
			if count < len(words) {
				fmt.Fprintln(output)
				continue
			}
			continue
		}
		// Process each character in the word
		for _, char := range word {
			asciiArt, err := functions.PrintChar(char, fontArray)
			if err != nil {
				log.Fatal("Error printing ascii", err)
			}
			// Concatenate the lines of the Ascii art for the character
			for i, line := range asciiArt {
				concatenatedLines[i] += line
			}
		}

		// Apply color if specified
		if *colorFlag != "" {
			colorParts := strings.SplitN(*colorFlag, " ", 2)
			if len(colorParts) == 2 {
				colorName := colorParts[0]
				colorCode, ok := colorMap[colorName]
				if !ok {
					log.Fatal("Invalid color name")
					return
				}
				concatenatedLines = functions.ColorSubstring(concatenatedLines, colorCode, substring, stringToColor)
			} else {
				log.Fatal("Usage: go run . [OPTION] [STRING]")
				return
			}
		}

		// Print Ascii art lines
		for _, line := range concatenatedLines {
			fmt.Fprintln(output, line)
		}
		fmt.Fprintln(output)
	}
}
