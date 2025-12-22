package main

import (
	"fmt"
)

func viewMeals(cfg *config, args ...string) error {
	for _, meal := range cfg.storedMeals {
		fmt.Println(meal)
	}
	fmt.Println()

	if len(cfg.storedMeals) == 0 {
		fmt.Println("No meals stored.")
		return nil
	}

	return nil
}
