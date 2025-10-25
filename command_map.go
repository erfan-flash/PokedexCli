package main

import (
	"errors"
	"fmt"
	"log"
)

func callbackMap(cfg *config) error {
	resp , err := cfg.pokeapiClient.LocationAreaList(cfg.nextLocationArea)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas: ")
	for _, area := range resp.Results{
		fmt.Printf("- %s\n", area.Name)
	}
	cfg.nextLocationArea = resp.Next
	cfg.prevLocationAtrea= resp.Previous
	return  nil
}
func callbackMapb(cfg *config) error {
	if cfg.prevLocationAtrea == nil{
		return errors.New("first page")
	}
	resp , err := cfg.pokeapiClient.LocationAreaList(cfg.prevLocationAtrea)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas: ")
	for _, area := range resp.Results{
		fmt.Printf("- %s\n", area.Name)
	}
	cfg.nextLocationArea = resp.Next
	cfg.prevLocationAtrea= resp.Previous
	return  nil
}