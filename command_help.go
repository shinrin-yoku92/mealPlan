package main

import (
	"fmt"
)

func showHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to mealPlan!")
	fmt.Println("Available commands:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
