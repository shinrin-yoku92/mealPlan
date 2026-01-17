package main

import (
	"fmt"
	"strings"
)

func removeMeal(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("Please provide a meal name")
	}

	mealName := strings.Join(args, " ")

	for i, meal := range cfg.storedMeals {
		if meal.Name == mealName {
			cfg.storedMeals = append(cfg.storedMeals[:i], cfg.storedMeals[i+1:]...)
			fmt.Printf("Meal '%s' removed successfully\n", mealName)
			return nil
		}
	}

	return fmt.Errorf("Meal '%s' not found", mealName)
}
