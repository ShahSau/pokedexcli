package main

type config struct {
	pokeapiClient           Client
	nextLocationAreaUrl     *string
	previousLocationAreaUrl *string
}

func main() {

	cfg := config{
		pokeapiClient: NewClient(),
	}

	startRepl(&cfg)
	// pokeapi := NewClient()
	// res, err := pokeapi.GetLocationAreas()

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(res)

}
