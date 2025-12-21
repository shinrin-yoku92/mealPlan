package main

import (
	"fmt"
	"strings"
)

func addMeal(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("please provide a meal name")
	}

	mealName := strings.Join(args, " ")

	cfg.storedMeals = append(cfg.storedMeals, mealName)
	fmt.Printf("Meal '%s' added successfully.\n", mealName)

	return cfg.Save(dataFile)
}
