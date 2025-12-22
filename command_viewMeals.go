package main

import (
	"fmt"
)

func viewMeals(cfg *config, args ...string) error {
	for _, meal := range cfg.storedMeals {
		fmt.Println(meal)
	}
	fmt.Println()
	
	return nil
}
