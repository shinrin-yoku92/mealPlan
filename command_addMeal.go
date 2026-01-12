package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func addMeal(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("please provide a meal name")
	}

	mealName := strings.Join(args, " ")

	for _, meals := range cfg.storedMeals {
		if meals.Name == mealName {
			fmt.Printf("'%s' already exists.\n", mealName)
			return nil
		}
	}

	reader := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter a description for '%s': ", mealName)
	reader.Scan()
	description := strings.TrimSpace(reader.Text())

	newMeal := Meal{
		Name: mealName,
	}
	if description != "" {
		newMeal.Description = description
	}

	cfg.storedMeals = append(cfg.storedMeals, newMeal)
	fmt.Printf("'%s' added successfully.\n", mealName)

	return cfg.Save(dataFile)
}
