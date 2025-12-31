package main

type config struct {
	Next			*string
	Previous	*string
	Pokemon 	map[string]pokemon
}
type pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

type LocationArea struct {
	Next						*string  `json:"next"`
	Previous 				*string  `json:"previous"`
	Results					[]struct {
		Name			string	 `json:"name"`
	} `json:"results"`
}

type ExploreConfig struct {
	Pokemon_encounters	[]struct {
		Pokemon struct {
			Name	string `json:"name"`
		} `json:"pokemon"`
	}	`json:"pokemon_encounters"`
}


