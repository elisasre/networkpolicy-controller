package controller

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func makeConfig(path string) (*Config, error) {
	config := &Config{}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
