package main

func main() {
	cfg := &config{
		storedMeals: map[string]string{},
	}

	startRepl(cfg)

}
