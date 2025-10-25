package main

import "fmt"

func callbackHelp(cfg *config) error{
	fmt.Println("Welcome to Pokedex Help menu!")
	fmt.Println("Available commands: ")
	availabe_commands := getCommands()
	for _ , cmd :=  range availabe_commands{
		fmt.Printf("- %s: %s \n", cmd.name, cmd.description)
		fmt.Println("")
	}
	return  nil
}