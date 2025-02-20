package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"Pokedex/internal/pokeapi"
)
type config struct {
	pokeapiClient *pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtpokemon map[string]pokeapi.Pokemon
}


func startrepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		//first word is the command
		command := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		// calling the command if it exists
		cmd, exists := getCommands()[command]
		if exists {
			err := cmd.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown cmd")
			continue
		}
	}
}


func cleanInput(text string) []string {
	s := strings.ToLower(text)
	result := strings.Split(s," ")
	return result
}

type cliCommand struct {
	name 			string
	description 	string
	callback 		func(*config, ...string) error
}

func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
		"exit": {
		name: 			"exit",
		description: 	"Exit the Pokedex",
		callback: 		commandExit,
		},
		"help": {
		name: 			"help",
		description: 	"Displays a help message",
		callback: 		commandHelp,
		},
		"map": {
		name: 			"map",
		description: 	"Displaying 20 areas/locations or forward to the next 20",
		callback: 		commandMapf,
		},
		"mapb": {
		name: 			"mapb",
		description: 	"Get the previous page of 20 locations",
		callback: 		commandMapb,
		},
		"explore": {
		name: 			"explore <area>",
		description: 	"Lets you explore which pokemon is found in a certain area",
		callback: 		commandExplore,
		},
		"catch": {
		name: 			"catch <pokemon>",
		description: 	"Attempt to catch a pokemon",
		callback: 		commandCatch,
		},
		"inspect": {
		name:			"inspect <pokemon>",
		description:	"Inspect the pokemon, you can only inspect pokemons you have caught",
		callback:		commandInspect,
		},
		"pokedex": {
		name: 			"pokedex",
		description:	"Show a list of all the pokemon you've captured",
		callback:		commandPokedex,
		},
	}
}