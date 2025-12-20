package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type command struct {
	name        string
	description string
	execute     func(*config, ...string) error
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("meal-tracker > ")
		reader.Scan()

		input := cleanInput(reader.Text())
		if len(input) == 0 {
			continue
		}

		commands := getCommands()
		cmd, exists := commands[input[0]]
		if exists {
			err := cmd.execute(cfg, input[1:]...)
			if err != nil {
				fmt.Println("Error executing command:", err)
			}
		} else {
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
			continue
		}

	}
}

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}

func getCommands() map[string]command {
	return map[string]command{
		"add": {
			name:        "add",
			description: "Add a meal with its description",
			execute:     addMeal,
		},
		"view": {
			name:        "view",
			description: "View all stored meals",
			execute:     viewMeals,
		},
		"plan": {
			name:        "plan",
			description: "Generate a meal plan",
			execute:     planMeals,
		},
		"remove": {
			name:        "remove",
			description: "Remove a meal by name",
			execute:     removeMeal,
		},
		"help": {
			name:        "help",
			description: "Show available commands",
			execute:     showHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the meal tracker",
			execute:     commandExit,
		},
	}
}
