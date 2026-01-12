package main

import (
	"fmt"
)

func viewMeals(cfg *config, args ...string) error {
	if len(cfg.storedMeals) == 0 {
		fmt.Println("No meals stored.")
		return nil
	}

	for _, meal := range cfg.storedMeals {
		if meal.Description == "" {
			fmt.Printf("%s\n", meal.Name)
		} else {
			fmt.Printf("%s - %s\n", meal.Name, meal.Description)
		}
	}
	fmt.Println()

	return nil
}
