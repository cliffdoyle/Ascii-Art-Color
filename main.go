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
	// Define the flags
	outputFlag := flag.String("output", "", "Specify file type")
	colorFlag := flag.String("color", "", "Specify color and substring to be colored in the format <color> <substring>")

	// Parse the flags
	flag.Parse()

	// Get non-flag arguments
	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("Usage: go run . [options] [String] [Font_File]")
		return
	}
	inputString := args[0]
	var fontFile string

	// Adjust to correctly handle the font file argument
	if len(args) == 2 && !strings.Contains(args[1], " ") {
		fontFile = args[1]
	} else {
		fontFile = "standard"
	}

	// Load the FontData
	fontArray, err := functions.Fonts(fontFile)
	if err != nil {
		fmt.Println("Error loading font:", err)
		return
	}

	// Prepare the inputString
	inputString = strings.ReplaceAll(inputString, "\\n", "\n")
	words := strings.Split(inputString, "\n")

	// Determine the output destination
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

	// Extract color and substring
	var colorCode string
	var substring string
	var ok bool
	if *colorFlag != "" {
		colorParts := strings.SplitN(*colorFlag, " ", 2)
		if len(colorParts) < 2 {
			log.Fatal("Usage: go run . --color=<color> <substring> [string] [Font_File]")
			return
		}
		colorName := colorParts[0]
		substring = strings.Join(colorParts[1:], " ")

		colorCode, ok = colorMap[colorName]
		if !ok {
			log.Fatal("Invalid color name")
			return
		}
	}

	// Process each word
	for _, word := range words {
		concatenatedLines := make([]string, 8)
		if word == "" {
			fmt.Fprintln(output)
			continue
		}
		// Process each character in the word
		for _, char := range word {
			asciiArt, err := functions.PrintChar(char, fontArray)
			if err != nil {
				log.Fatal("Error printing ascii:", err)
			}
			// Concatenate the lines of the Ascii art for the character
			for i, line := range asciiArt {
				concatenatedLines[i] += line
			}
		}
		// Apply color if specified
		if *colorFlag != "" {
			concatenatedLines = functions.ColorSubstring(concatenatedLines, colorCode, substring)
		}
		// Print Ascii art lines
		for _, line := range concatenatedLines {
			fmt.Fprintln(output, line)
		}
		fmt.Fprintln(output)
	}
}
