package main

import (
	"fmt"
	"slices"
	"strings"
)

func addMeal(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("please provide a meal name")
	}

	mealName := strings.Join(args, " ")

	if slices.Contains(cfg.storedMeals, mealName) {
		fmt.Printf("'%s' already exists.\n", mealName)
		return nil
	}

	cfg.storedMeals = append(cfg.storedMeals, mealName)
	fmt.Printf("'%s' added successfully.\n", mealName)

	return cfg.Save(dataFile)
}
