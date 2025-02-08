package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var commands map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next string
	Previous string
}

type locationResponse struct {
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
	Next string `json:"next"`
	Previous string `json:"previous"`
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
				err := command.callback(cfg)
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

func commandExit(cfg *config) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp(cfg *config) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage: \n\n\n")
    for _, cmd := range commands {
        fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    }
    return nil
}

func commandMap(cfg *config) error {
	if cfg.Next == "" {
		return fmt.Errorf("no more locations to display")
	}
	res, err := http.Get(cfg.Next)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var loc locationResponse
	if err := json.NewDecoder(res.Body).Decode(&loc); err != nil {
		return err
	}
	for _, loc := range loc.Results {
		fmt.Println(loc.Name)
	}
	cfg.Next = loc.Next
	cfg.Previous = loc.Previous
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == "" {
		return fmt.Errorf("no more locations to display")
	}
	res, err := http.Get(cfg.Previous)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var loc locationResponse
	if err := json.NewDecoder(res.Body).Decode(&loc); err != nil {
		return err
	}
	for _, loc := range loc.Results {
		fmt.Println(loc.Name)
	}
	cfg.Next = loc.Next
	cfg.Previous = loc.Previous
	return nil
}