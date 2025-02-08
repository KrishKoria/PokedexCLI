package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commands map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to PokedexCLI! Type 'exit' to quit.")
	for {
		fmt.Print("PokeDex > ")
		if !scanner.Scan() {
            break
        }
		text := scanner.Text()
		words := cleanInput(text)
		if len(words) > 0 {
			if command, ok := commands[words[0]]; ok {
				err := command.callback()
				if err != nil {
					fmt.Println("Error executing command:", err)
				}
			} else {
				fmt.Println("Unknown command:", words[0])
			}
		}
	}
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
    text = strings.ToLower(text)
    text = strings.ReplaceAll(text, ",", " ")
	words := strings.Fields(text)
	return words
}

func commandExit() error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp() error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage: \n\n\n")
    for _, cmd := range commands {
        fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    }
    return nil
}