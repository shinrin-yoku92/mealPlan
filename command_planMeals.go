package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func planMeals(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please specify the number of meals to plan.")
	}

	numMeals, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("Invalid number of meals: %s", err)
	}

	if numMeals <= 0 {
		return fmt.Errorf("Number of meals must be greater than 0.")
	}

	if len(cfg.storedMeals) == 0 {
		return fmt.Errorf("No meals available in the storage.")
	}

	if numMeals > len(cfg.storedMeals) {
		numMeals = len(cfg.storedMeals)
	}

	indices := rand.Perm(len(cfg.storedMeals))

	fmt.Printf("\nMeal plan (%d meals):\n", numMeals)
	fmt.Println(strings.Repeat("-", 20))
	for i := range numMeals {
		meal := cfg.storedMeals[indices[i]]
		fmt.Printf("%d. %s", i+1, meal.Name)
		if meal.Description != "" {
			fmt.Printf(" - %s", meal.Description)
		}
		fmt.Println()
	}
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println()

	return nil
}
