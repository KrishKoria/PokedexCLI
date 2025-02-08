package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
    text = strings.ReplaceAll(text, ",", " ")
	words := strings.Fields(text)
	return words
}