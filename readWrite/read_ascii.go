package readwrite

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var files = map[string]bool{
	"banners/shadow.txt":     true,
	"banners/standard.txt":   true,
	"banners/thinkertoy.txt": true,
}

func ValidateFileName(file string) bool {
	_, ok := files[file]
	return ok
}

func ReadAscii(fileName string) (map[int][]string, error) {
	if !ValidateFileName(fileName) {
		return nil, fmt.Errorf("unsupported file name: %s", fileName)
	}

	if !strings.HasSuffix(fileName, ".txt") {
		return nil, fmt.Errorf("unsupported file format: %s", fileName)
	}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	bannerMap := make(map[int][]string)
	key := 32
	lineCount := 0
	currentart := []string{}

	for scanner.Scan() {
		lines := scanner.Text()

		if lines != "" {
			currentart = append(currentart, lines)
			lineCount++
		}

		if lineCount == 8 {
			bannerMap[key] = currentart
			key++
			currentart = []string{}
			lineCount = 0
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return bannerMap, nil
}
