package main

type config struct {
	Next			*string
	Previous	*string
}

type LocationArea struct {
	Next						*string  `json:"next"`
	Previous 				*string  `json:"previous"`
	Results					[]struct {
		Name			string	 `json:"name"`
	} `json:"results"`
}


