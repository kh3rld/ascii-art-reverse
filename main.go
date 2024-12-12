package main

import (
	"fmt"

	asciiArt "github.com/Doreen-Onyango/ascii-art-reverse/readWrite"
	utils "github.com/Doreen-Onyango/ascii-art-reverse/utils"
)

func main() {
	if !utils.CheckArgs() {
		fmt.Printf("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>\n")
		return
	}

	standardFile := "banners/standard.txt"
	shadowFile := "banners/shadow.txt"
	thinkertoyFile := "banners/thinkertoy.txt"

	utils.ParseFlag()
	sampleFile := *utils.ReversePtr

	if sampleFile == "" {
		asciiArt.GenerateArt()
	} else {
		bannerFiles := []string{standardFile, shadowFile, thinkertoyFile}
		var decodedString string
		var success bool

		for _, bannerFile := range bannerFiles {
			artToChar := utils.ParseBannerFile(bannerFile)
			lines, err := utils.ParseFile(sampleFile)
			if err != nil {
				fmt.Println(err)
				return
			}
			decodedString = utils.DecodeFile(lines, artToChar)
			if decodedString != "" {
				success = true
				break
			}
		}
		if success {
			fmt.Print(decodedString)
		} else {
			fmt.Println("Error: Could not decode the ASCII art using any of the banner files.")
		}
	}
}
