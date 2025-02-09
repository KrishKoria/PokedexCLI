package main

import (
	"bufio"
	"fmt"
	"math/rand"
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
	Pokedex map[string]pokeapi.Pokemon
}


func main() {
	cfg := &config{
		Next: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		Previous: "",
		Pokedex: make(map[string]pokeapi.Pokemon),
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
		"catch": {
			name: "catch",
			description: "Catch a Pokemon",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Inspect a Pokemon",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "View your Pokedex",
			callback: commandPokedex,
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

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a Pokemon name")
	}

	pokemonName := args[0]
	pokemon, err := pokeapi.FetchPokemon(pokemonName)
	if err != nil {
		return err
	}
	
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	
	baseExperience := pokemon.BaseExperience
    catchProbability := 100 - baseExperience/2
    if catchProbability < 10 {
        catchProbability = 10 
    }
	
	chance := rand.Intn(100)
    if chance < catchProbability{
        fmt.Printf("%s was caught!\n", pokemonName)
		fmt.Println("You may now inspect it with the inspect command.")
        cfg.Pokedex[pokemonName] = *pokemon
    } else {
        fmt.Printf("%s escaped!\n", pokemonName)
    }

	return nil
}

func commandInspect(cfg *config, args []string) error {
    if len(args) == 0 {
        return fmt.Errorf("please provide a Pokemon name")
    }

    pokemonName := args[0]
    pokemon, ok := cfg.Pokedex[pokemonName]
    if !ok {
        fmt.Printf("You have not caught that Pokemon\n")
        return nil
    }

    fmt.Printf("Name: %s\n", pokemon.Name)
    fmt.Printf("Height: %d\n", pokemon.Height)
    fmt.Printf("Weight: %d\n", pokemon.Weight)
    fmt.Println("Stats:")
    for _, stat := range pokemon.Stats {
        fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
    }
    fmt.Println("Types:")
    for _, t := range pokemon.Types {
        fmt.Printf("  - %s\n", t.Type.Name)
    }

    return nil
}

func commandPokedex(cfg *config, args []string) error {
    if len(cfg.Pokedex) == 0 {
        fmt.Println("Your Pokedex is empty.")
        return nil
    }

    fmt.Println("Your Pokedex:")
    for name := range cfg.Pokedex {
        fmt.Printf(" - %s\n", name)
    }
    return nil
}
