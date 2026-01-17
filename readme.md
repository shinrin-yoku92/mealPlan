# mealPlan

A simple **CLI app** that randomly selects meals for a specified number of days.

This tool helps you plan meals by letting you manage a list of meals and then generate a random meal plan for a given number of days.

## 🚀 Features

- Add meals to your list  
- Remove meals  
- View all saved meals  
- Generate a random meal plan for *X* days  
- Interactive REPL (Read–Eval–Print Loop) for quick command use  

## 🛠️ Getting Started

### Prerequisites

You’ll need Go installed on your system (Go 1.18+ recommended).

Check installation with:

```sh
go version
```

### Clone the repo

```sh
git clone https://github.com/shinrin-yoku92/mealPlan.git
cd mealPlan
```

### Build

```sh
go build -o mealplan
```

This produces a `mealplan` executable in your working directory.

### Run

Start the CLI:

```sh
./mealplan
```

Once launched, use the built-in commands to manage and generate meal plans.

## 📋 Usage

Below are common commands you can use in the CLI:

| Command             | Description                           |
| ------------------- | ------------------------------------- |
| `add <name>`    | Adds a meal to your list              |
| `remove <name>` | Removes a meal by name                |
| `view`         | Lists all current meals               |
| `plan <days>`  | Generates a random meal plan for days |
| `help`              | Lists available commands              |
| `exit`              | Quit the CLI                          |

The CLI supports interactive input. You can type these commands once the program is running.

## 🧪 Examples

Add a meal:

```sh
> add Spaghetti
Meal added: Spaghetti
```

Generate a random 5-day plan:

```sh
> plan 5
Day 1: Tacos
Day 2: Spaghetti
...
```

## 📄 License

This project is open source and available under the terms of the **MIT License**.
```
