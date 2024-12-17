package utils

import "fmt"

func PrintUsage() {
	fmt.Print("Usage: go run . [OPTION] [STRING] [BANNER]\n\n")
	fmt.Println("EX: go run . --reverse=<fileName.txt>")
}
