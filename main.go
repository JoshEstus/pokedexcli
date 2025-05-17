package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			cleaned := cleanInput(userInput)
			fmt.Printf("Your command was: %s\n", cleaned[0])
		}

	}
}

func cleanInput(text string) []string {
	// trim starting ane ending whitespace
	cleanedText := strings.TrimSpace(text)
	re := regexp.MustCompile(`\s+`)
	// clean internal extra spaces
	cleanedText = re.ReplaceAllString(cleanedText, " ")
	cleanedText = strings.ToLower(cleanedText)
	cleanedInput := strings.Split(cleanedText, " ")

	return cleanedInput
}
