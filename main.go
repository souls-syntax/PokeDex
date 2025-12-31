package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
	"github.com/souls-syntax/PokeDex/internal/pokecache"
	"time"
	

)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	Commands := GetCommands()
	
	cfg := &config {
			Next: nil,
			Previous: nil,
			Pokemon: make(map[string]pokemon),
		}
	
	cache := pokecache.NewCache(5*time.Second)
	
	for {
		fmt.Printf("Pokedex >")
	
		scanner.Scan()
		
		text := scanner.Text()
		
		var cleanString []string
		
		cleanString = cleanInput(text)
		
		if len(cleanString) == 0{
			continue
		}	
		
		cmd,ok := Commands[cleanString[0]]
		
		if !ok{
			fmt.Println("Unknown Command")
		} else {
			if err := cmd.Callback(cfg,cache,cleanString[1:]); err != nil {
				fmt.Print(err)
			}
		}
	}
}

func cleanInput(text string) []string{
	return strings.Fields(strings.ToLower(text))	
}
