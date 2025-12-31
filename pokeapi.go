package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
	"math/rand"
	"github.com/souls-syntax/PokeDex/internal/pokecache"
)

func CommandMap(cfg *config, c *pokecache.Cache, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	var apiurl string
	if cfg.Next == nil {
		apiurl = url
	} else {
		apiurl = *cfg.Next
	}
	dec, err := getCacheLocation(apiurl, c)
	if err != nil {
		return fmt.Errorf("Error in getCacheLocation")
	}
	var DataLocation LocationArea
	if err := json.Unmarshal(dec,&DataLocation); err != nil {
		return fmt.Errorf("Error : Problem with Unmarshaling the io")
	}
	cfg.Next = DataLocation.Next
	cfg.Previous = DataLocation.Previous
	
	for _,addr := range DataLocation.Results {
		fmt.Println(addr.Name)
	}
	return nil
}


func CommandMapBack(cfg *config, c *pokecache.Cache, args []string) error {
	var apiurl string
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	} else {
		apiurl = *cfg.Previous
	}
	dec, err := getCacheLocation(apiurl, c)
	if err != nil {
		return fmt.Errorf("Error in getCacheLocation")
	}
	var DataLocation LocationArea

	if err := json.Unmarshal(dec,&DataLocation); err != nil {
		return fmt.Errorf("Error : Problem with Unmarshaling the io")
	}
	cfg.Next = DataLocation.Next
	cfg.Previous = DataLocation.Previous
	
	for _,addr := range DataLocation.Results {
		fmt.Println(addr.Name)
	}
	return nil
}


func getCacheLocation(url string, c *pokecache.Cache) ([]byte, error) {
	val, ok := c.Get(url)
	if ok {
		return val, nil
	}	
	req,err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil,fmt.Errorf("Error : NewRequest")
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil,fmt.Errorf("Error : client")
	}
	
	defer res.Body.Close()

	dec,_ := io.ReadAll(res.Body)
	c.Add(url,dec)

	return dec, nil
}

func CommandExplore(cfg *config, c *pokecache.Cache, args []string) error {
	if len(args) == 0 {
		fmt.Println("Please specify the area to explore")
		return nil
	}
	city := args[0]
	
	url := "https://pokeapi.co/api/v2/location-area/" + city
	val,_ := getCacheLocation(url,c)

	var exploredata ExploreConfig

	if err := json.Unmarshal(val,&exploredata); err != nil {
		return fmt.Errorf("Error: CommandExplore")
	}
	fmt.Printf("Exploring %v...\n", city)
	fmt.Println("Found Pokemon:")
	for _,obj := range exploredata.Pokemon_encounters {
		fmt.Printf(" - %v\n",obj.Pokemon.Name)
	}
	return nil
}	
	
func CommandCatch(cfg *config, c *pokecache.Cache, args []string) error {
	if len(args) == 0 {
		fmt.Printf("Specify the pokemon\n")
		return nil
	}
	pokemonName := args[0]
	url := "https://pokeapi.co/api/v2/pokemon/"+pokemonName
	val,err := getCacheLocation(url, c)
	if err != nil {
		return fmt.Errorf("Error: CommandCatch")
	}
	var poke pokemon
	if err := json.Unmarshal(val,&poke); err != nil {
		return fmt.Errorf("Error : CommandCatch Unmarshal")
	}
	fmt.Printf("Throwing a Pokeball at %v...\n",pokemonName)
	difficulty := poke.BaseExperience
	res := rand.Intn(difficulty)
	if res > 40 {
		fmt.Printf("%v escaped!\n",pokemonName)
		return nil
	} else {
		fmt.Printf("%v was caught!\n",pokemonName)
		cfg.Pokemon[pokemonName] = poke
		return nil
	}
}

func CommandInspect(cfg *config, c *pokecache.Cache, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, ok := cfg.Pokemon[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}

func CommandPokedex(cfg *config, c *pokecache.Cache, args []string) error {
	fmt.Println("Your Pokedex:")
	for pokemon,_ := range cfg.Pokemon {
		fmt.Printf(" - %s\n",pokemon)
	}
	return nil
}


