package main

import (
	"fmt"
	"os"
	"github.com/souls-syntax/PokeDex/internal/pokecache"
)

type CliCommand struct {
	Name					string
	Description 	string
	Callback			func(*config, *pokecache.Cache) error
}

func GetCommands() map[string]CliCommand { 
	return map[string]CliCommand{
		"exit": {
			Name:					"exit",
			Description:	"Exit the Pokedex",
			Callback:			CommandExit,
		},
		"help": {
			Name:					"help",
			Description:	"Display a help message",
			Callback:			CommandHelp,
		},
		"map": {
			Name:					"map",
			Description:	"Display the name next of 20 Location",
			Callback:			CommandMap,
		},
		"mapb": {
			Name:					"mapb",
			Description:	"Display the name previous of 20 Location",
			Callback:			CommandMapBack,
		},
	}
}
func CommandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	Commands := GetCommands()
	for _,cmd := range Commands {
		fmt.Printf("%v: %v\n",cmd.Name,cmd.Description)
	}
	return nil
}

