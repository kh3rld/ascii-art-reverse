package utils

import (
	"flag"
	"os"
	"strings"
)

// CheckArgs checks the validity of the arguments passed.
func CheckArgs() bool {
	switch len(os.Args) {
	case 1:
		return false
	case 2:
		return !(strings.HasPrefix(os.Args[1], "-") && !strings.HasPrefix(os.Args[1], "--reverse="))
	case 3:
		if strings.HasPrefix(os.Args[1], "-") {
			return false
		}
		banner := strings.ToLower(os.Args[2])
		return banner == "standard" || banner == "shadow" || banner == "thinkertoy"
	}
	return false
}

// DecodeFile decodes the ASCII art from the provided lines using the art-to-character map.
func DecodeFile(lines []string, artToChar map[string]string) string {
	var result strings.Builder
	for i := 0; i < len(lines); {
		if lines[i] == "" {
			result.WriteString("\n")
			i++
			continue
		}
		if i+7 < len(lines) {
			decoded := decode(lines[i:i+8], artToChar)
			if decoded == "" {
				return ""
			}
			result.WriteString(decoded + "\n")
			i += 8
		} else {
			i++
		}
	}
	return result.String()
}

// decodeChunk decodes a single chunk of ASCII art using the art-to-character map.
func decode(chunk []string, artToChar map[string]string) string {
	const artHeight = 8
	var decodedString strings.Builder

	sliceArt := func(chunk []string, start, width int) []string {
		art := make([]string, artHeight)
		for row := 0; row < artHeight; row++ {
			if start+width <= len(chunk[row]) {
				art[row] = chunk[row][start : start+width]
			} else {
				art[row] = chunk[row][start:]
			}
		}
		return art
	}

	for col := 0; col < len(chunk[0]); {
		found := false
		for width := 1; col+width <= len(chunk[0]); width++ {
			subArt := sliceArt(chunk, col, width)
			if val, exists := artToChar[strings.Join(subArt, "\n")]; exists {
				decodedString.WriteString(val)
				col += width
				found = true
				break
			}
		}
		if !found {
			col++
		}
	}

	return decodedString.String()
}

var ReversePtr = flag.String("reverse", "", "reverse the ascii art")

// ParseFlag parses command-line flags.
func ParseFlag() {
	flag.Parse()
}
