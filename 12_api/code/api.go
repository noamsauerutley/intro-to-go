package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Planet is a planet type
type Planet struct {
	Name       string `json:"name"`
	Terrain    string `json:"terrain"`
	Population string `json:"population"`
}

// Person is a person type
type Person struct {
	Name         string `json:"name"`
	HomeworldURL string `json:"homeworld"`
	Homeworld    Planet
}

// AllPeople is a collection of person types
type AllPeople struct {
	People []Person `json:"results"`
}

// BaseURL is the base endpoint for the star wars API
const BaseURL = "https://swapi.dev/api/"

func (p *Person) getHomeworld() {
	res, err := http.Get(p.HomeworldURL)
	if err != nil {
		log.Print("Error fetching homeworld", err)
	}

	var bytes []byte
	if bytes, err = ioutil.ReadAll(res.Body); err != nil {
		log.Print("Error reading response body", err)
	}

	json.Unmarshal(bytes, &p.Homeworld)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "People successfully retrieved")

	var people AllPeople

	res, err := http.Get(BaseURL + "people")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to retrieve people")
	}

	fmt.Println(res)

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to parse request body")
	}

	// fmt.Println(string(bytes))

	if err := json.Unmarshal(bytes, &people); err != nil {
		fmt.Println("Error parsing json", err)
	}

	for _, individual := range people.People {
		individual.getHomeworld()
		fmt.Println(individual)
	}

	fmt.Println("People: ", people)
}

func main() {
	http.HandleFunc("/people/", getPeople)
	fmt.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
