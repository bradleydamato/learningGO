package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"john", "bud"}},
	{Title: "Cool Hand Luke", Year: 1212, Color: true, Actors: []string{"zack", "kyle"}},
	{Title: "Bullitt", Year: 2020, Color: true, Actors: []string{"brad", "jack"}},
}

func main() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

}
