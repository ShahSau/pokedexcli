package main

import "time"

type config struct {
	pokeapiClient           Client
	nextLocationAreaUrl     *string
	previousLocationAreaUrl *string
}

func main() {

	cfg := config{
		pokeapiClient: NewClient(time.Hour),
	}

	startRepl(&cfg)
	// pokeapi := NewClient()
	// res, err := pokeapi.GetLocationAreas()

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(res)

}
