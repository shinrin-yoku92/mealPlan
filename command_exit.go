package main

import (
	"fmt"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Exiting mealPlan. Goodbye!")
	return fmt.Errorf("exit")
}
