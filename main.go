package main

import "github.com/erfan-flash/PokedexCli/internals/pokedeapi"

type config struct {
	pokeapiClient pokedeapi.Client
	nextLocationArea *string
	prevLocationAtrea *string
}

func main() {
	cfg := config{
		pokeapiClient: pokedeapi.NewClient(),
	}
	startRepl(&cfg)
}