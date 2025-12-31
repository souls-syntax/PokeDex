package main 

import (
	"fmt"
	"os"
	"github.com/souls-syntax/PokeDex/internal/pokecache"
)

type CliCommand struct {
	Name					string
	Description 	string
	Callback			func(*config, *pokecache.Cache, []string) error
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
		"explore": {
			Name:					"explore",
			Description:	"Display the pokemon available in the area",
			Callback:			CommandExplore,
		},
		"catch": {
			Name:					"catch",
			Description:	"Display the pokemon available in the area",
			Callback:			CommandCatch,
		},
		"inspect": {
			Name:					"inspect",
			Description:	"Know about the pokemon you have catched",
			Callback:			CommandInspect,
		},
		"pokedex": {
			Name:					"pokedex",
			Description:	"Know all the pokemon you have catched",
			Callback:			CommandPokedex,
		},
	}
}

func CommandExit(cfg *config, c *pokecache.Cache, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(cfg *config, c *pokecache.Cache, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	Commands := GetCommands()
	for _,cmd := range Commands {
		fmt.Printf("%v: %v\n",cmd.Name,cmd.Description)
	}
	return nil
}

