package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, command := range Commands {
		cliHelp := fmt.Sprintf("%s: %s", command.name, command.description)
		fmt.Println(cliHelp)
	}
	return nil
}
