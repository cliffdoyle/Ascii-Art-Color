package functions

import (
	"strings"
	"fmt"
)

// ColorSubstring colors the specified substring within the concatenatedLines using the given colorCode.
func ColorSubstring(concatenatedLines []string, color, substring string) []string {
	if substring == "" {
		//Color the whole string
		for i := range concatenatedLines {
			concatenatedLines[i] = fmt.Sprintf("\033[%sm%s\033[0m", color, concatenatedLines[i])
        
}
	} else {
		//color the specified substring
		for i := range concatenatedLines {
			concatenatedLines[i] = strings.ReplaceAll(concatenatedLines[i],substring,fmt.Sprintf("\033[%sm%s\033[0m", color, substring))
		}
	}
	return concatenatedLines
}
