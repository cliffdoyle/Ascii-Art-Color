package main

import (
	"ascii-art-color/functions"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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
	if len(os.Args) == 1 || len(os.Args) > 4 {
		return
	}

	color := flag.String("color", "", "Specify color")
	flag.Parse()
	args := flag.Args()
	var inputString string
	var substring string
	fontfile := "standard"

	switch len(args) {
	case 1:
		inputString = args[0]
	case 2:
		substring = args[0]
		inputString = args[1]
	case 3:
		substring = args[0]
		inputString = args[1]
		fontfile = args[2]

	}

	//Load fontdata
	fontArray, err := functions.Fonts(fontfile)
	if err != nil {
		log.Fatal(err)
	}

	//Process the color flag

	var colorCode string
	var ok bool
	if *color != "" {
		colorCode, ok = colorMap[*color]
		if !ok {
			fmt.Println("Invalid color")
			return
		}
	}

	//Process the inputString

	inputString = strings.ReplaceAll(inputString, "\\n", "\n")

	lines := strings.Split(inputString, "\n")
	count := 0
	for _, line := range lines {
		concatenatedLines := make([]string, 8)

		//Handle newline
		if line == "" {
			count++
			if count < len(lines) {
				fmt.Println()
				continue
			}
			continue
		}

		//Determine substring Destination
		var sublines []string
		if substring == "" {
			sublines = []string{line}
		} else {
			sublines = strings.Split(line, substring)
		}

		for i, subline := range sublines {

			for _, char := range subline {
				asciiArt, err := functions.PrintChar(char, fontArray)
				if err != nil {
					log.Fatal(err)
				}
				for j, artline := range asciiArt {
					concatenatedLines[j] += artline
				}
			}

			if substring != "" && i < len(sublines)-1 {
				coloredSub := functions.ColorSubstring(substring, fontArray, colorCode)
				for k, artline := range coloredSub {
					concatenatedLines[k] += artline
				}
			}
		}
		if substring == "" && colorCode != "" {
			concatenatedLines = functions.ColorSubstring(line, fontArray, colorCode)
		}

		for _, artlines := range concatenatedLines {
			fmt.Println(artlines)
		}

	}

}
