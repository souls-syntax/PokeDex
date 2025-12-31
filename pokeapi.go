package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
	"github.com/souls-syntax/PokeDex/internal/pokecache"
)

func CommandMap(cfg *config, c *pokecache.Cache) error {
	url := "https://pokeapi.co/api/v2/location-area"
	var apiurl string
	if cfg.Next == nil {
		apiurl = url
	} else {
		apiurl = *cfg.Next
	}
	dec, err := getCacheLocation(apiurl, c)
	if err != nil {
		return fmt.Errorf(err)
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


func CommandMapBack(cfg *config, c *pokecache.Cache) error {
	var apiurl string
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	} else {
		apiurl = *cfg.Previous
	}
	dec, err := getCacheLocation(apiurl, c)
	if err != nil {
		return fmt.Errorf(err)
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
		return fmt.Errorf("Error : NewRequest")
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error : client")
	}
	
	defer res.Body.Close()

	dec,_ := io.ReadAll(res.Body)
	c.Add(url,dec)

	return dec, nil
}
