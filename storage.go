package main

import (
	"encoding/json"
	"os"
)

type config struct {
	storedMeals map[string]string
}

func (cfg *config) Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // No file to load, start fresh
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&cfg.storedMeals)
}

func (cfg *config) Save(path string) error {
	tmp := path + ".tmp"
	b, err := json.MarshalIndent(cfg.storedMeals, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(tmp, b, 0644); err != nil {
		return err
	}

	return os.Rename(tmp, path)
}
