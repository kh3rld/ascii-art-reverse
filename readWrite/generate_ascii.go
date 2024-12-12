package readwrite

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Doreen-Onyango/ascii-art-reverse/checksum"
	"github.com/Doreen-Onyango/ascii-art-reverse/utils"
)

func GenerateArt() {
	if len(os.Args) < 2 {
		utils.PrintUsage()
		return
	}

	input := os.Args[1]
	if input == "\\n" {
		fmt.Println()
		return
	}
	filename := RetieveBannerFile()
	err := checksum.ValidateFileChecksum(filename)
	if err != nil {
		log.Printf("Error downloading or validating file: %v", err)
		return
	}
	pattern, err := ReadAscii(filename)
	if err != nil {
		fmt.Println("error loading banner map:", err)
		return
	}
	args := strings.ReplaceAll(os.Args[1], "\\n", "\n")
	args = strings.ReplaceAll(args, "\\t", "    ")
	lines := strings.Split(args, "\n")

	for _, line := range lines {
		RenderBannerLine(line, pattern)
	}
}

func RetieveBannerFile() string {
	if len(os.Args) == 3 {
		switch os.Args[2] {
		case "standard":
			return "banners/standard.txt"
		case "shadow":
			return "banners/shadow.txt"
		case "thinkertoy":
			return "banners/thinkertoy.txt"
		default:
			return ""
		}
	}
	if len(os.Args) == 2 {
		return "banners/standard.txt"
	}
	return ""
}

// Print the banner for a line of text
func RenderBannerLine(line string, bannerMap map[int][]string) {
	if line == "" {
		fmt.Println()
		return
	}

	output := make([]string, 8)
	for _, char := range line {
		banner, exists := bannerMap[int(char)]
		if !exists {
			fmt.Printf("Character %c not found in banner map\n", char)
			continue
		}

		for i := 0; i < 8; i++ {
			output[i] += banner[i]
		}
	}

	for _, outLine := range output {
		fmt.Println(outLine)
	}
}
