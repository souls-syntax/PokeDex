package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)



func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Commands := GetCommands()
	cfg := &config {
			Next: nil,
			Previous: nil,
		}
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
			if err := cmd.Callback(cfg); err != nil {
				fmt.Print(err)
			}
		}
	}
}

func cleanInput(text string) []string{
	return strings.Fields(strings.ToLower(text))	
}


