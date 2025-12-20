package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Exiting mealPlan. Goodbye!")
	os.Exit(0)
	return nil
}
