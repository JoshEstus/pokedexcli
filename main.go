package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Declare a global map
var Commands map[string]cliCommand

func initializeMap() {
	Commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "display a help message",
			callback:    commandHelp,
		},
	}
}

func main() {
	initializeMap()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			cleaned := cleanInput(userInput)

			command, ok := Commands[cleaned[0]]

			if !ok {
				fmt.Println("Unknown command")
				continue
			}
			command.callback()
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
