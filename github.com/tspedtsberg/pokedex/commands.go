package main

import (
	"fmt"
	"os"
	"errors"
	"math/rand"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmds := range getCommands() {
		fmt.Printf("%s: %s\n", cmds.name, cmds.description)
	}
	return nil
}

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMapf(cfg *config, args ...string) error {
	sliceLocations, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = sliceLocations.Next
	cfg.prevLocationsURL = sliceLocations.Previous

	fmt.Println("List of 20 locations:")
	fmt.Println("--------------------")
	fmt.Println("")
	for _, loc := range sliceLocations.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println("")
	fmt.Println("---------- To see next page ---------")
	fmt.Println("Type 'map' again to see the next 20 locations")
	fmt.Println("")
	fmt.Println("---------- To see prev page ---------")
	fmt.Println("Type 'mapb' to see the prev 20 locations")
	return nil
}	

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("You're on the first page")
	}
	sliceLocations, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = sliceLocations.Next
	cfg.prevLocationsURL = sliceLocations.Previous

	for _, loc := range sliceLocations.Results {
		fmt.Println(loc.Name)
	}
	return nil
	}
	
func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location to explore")
	}

	nameOfLocation := args[0]
	location, err := cfg.pokeapiClient.ListPokemons(nameOfLocation)
	if err != nil {
		return fmt.Errorf("You must provide a valid area to explore")
	}

	for _, enc := range location.PokemonEncounters {
		fmt.Println(enc.Pokemon.Name)
	}

	return nil

}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon to catch")
	}

	nameOfPokemon := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(nameOfPokemon)
	if err != nil {
		return fmt.Errorf("You must provide a valid pokemon to catch")
	}
	//creating randomness
	chance := rand.Intn(pokemon.BaseExperience)


	fmt.Printf("Throwing a Pokeball at %s...", pokemon.Name)
	if chance > 540 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
		
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)
	cfg.caughtpokemon[pokemon.Name] = pokemon	

	return nil
}


func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon to inspect")
	}

	nameOfPokemon := args[0]

	pokemon, ok := cfg.caughtpokemon[nameOfPokemon]
	if !ok {
		return errors.New("you have not captured that pokemon yet")
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, stat := range pokemon.Types {
		fmt.Printf(" - %s\n", stat.Type.Name)
	}

	return nil
}


func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, v := range cfg.caughtpokemon {
		fmt.Printf(" - %s\n", v.Name)
	}

	return nil
}