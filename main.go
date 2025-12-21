package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var dataFile string

func init() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Printf("Error determining executable path: %v\n", err)
		os.Exit(1)
	}
	dataFile = filepath.Join(filepath.Dir(exePath), "meals.json")
}

func main() {
	cfg := &config{
		storedMeals: []string{},
	}

	if err := cfg.Load(dataFile); err != nil {
		fmt.Printf("Error loading meals: %v\n", err)
		return
	}

	defer func() {
		if err := cfg.Save(dataFile); err != nil {
			fmt.Printf("Error saving meals: %v\n", err)
		}
	}()

	startRepl(cfg)

}
