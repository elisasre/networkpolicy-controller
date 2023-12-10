package controller

import (
	"encoding/json"
	"log"
	"os"
)

func makeConfig(path string) (*Config, error) {
	config := &Config{}
	log.Printf("Loading config from path %s", path)
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// Contains is checking does array contain single word.
func Contains(array []string, word string) bool {
	for _, item := range array {
		if item == word {
			return true
		}
	}
	return false
}
