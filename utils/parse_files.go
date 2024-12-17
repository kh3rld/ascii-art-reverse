package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ParseBannerFile reads a banner file into a map of ASCII art to characters.
func ParseBannerFile(filename string) map[string]string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	artToChar := make(map[string]string)
	var art []string
	charCode := 32

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(art) == 8 {
				artToChar[strings.Join(art, "\n")] = string(rune(charCode))
				art = []string{}
				charCode++
			}
		} else {
			art = append(art, line)
		}
	}

	if len(art) == 8 {
		artToChar[strings.Join(art, "\n")] = string(rune(charCode))
	}

	return artToChar
}

// ParseFile reads a file and returns its lines as a slice of strings.
func ParseFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return lines, nil
}
