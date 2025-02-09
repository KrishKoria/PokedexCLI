package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeapi "github.com/KrishKoria/PokeApi"
)

var commands map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	Next string
	Previous string
}



func main() {
	cfg := &config{
		Next: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		Previous: "",
	}
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
		"map": {
			name:        "map",
			description: "Displays the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous map",
			callback:    commandMapb,
		},
		"explore" : {
			name: "explore",
			description: "Explore the map",
			callback: commandExplore,
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
				err := command.callback(cfg, words[1:])
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

func commandExit(cfg *config, args []string) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp(cfg *config, args []string) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Print("Usage: \n\n")
    for _, cmd := range commands {
        fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    }
    return nil
}

func commandMap(cfg *config, args []string) error {
	if cfg.Next == "" {
		return fmt.Errorf("no more locations to display")
	}
	loc, err := pokeapi.FetchLocations(cfg.Next)
	if err != nil {
		return err
	}
	for _, loc := range loc.Results {
		fmt.Println(loc.Name)
	}
	cfg.Next = loc.Next
	cfg.Previous = loc.Previous
	return nil
}

func commandMapb(cfg *config, args []string) error {
	if cfg.Previous == "" {
		return fmt.Errorf("no more locations to display")
	}
	loc, err := pokeapi.FetchLocations(cfg.Previous)
	if err != nil {
		return err
	}
	for _, loc := range loc.Results {
		fmt.Println(loc.Name)
	}
	cfg.Next = loc.Next
	cfg.Previous = loc.Previous
	return nil
}

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
        return fmt.Errorf("please provide a location area name")
    }
	locationName := args[0]
	locationArea, err := pokeapi.FetchLocationArea(locationName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", locationName)
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}
	return nil
}