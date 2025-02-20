package main

import (
	"time"
	"Pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtpokemon: map[string]pokeapi.Pokemon{},
	}

	startrepl(cfg)
}
	
