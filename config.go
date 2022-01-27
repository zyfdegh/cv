package main

import (
	"encoding/json"
	"os"
)

type Style struct {
	Mode     Mode   `json:"mode"`
	Template string `json:"template"`
}

type Config struct {
	Data  map[string]interface{} `json:"data"`
	Style Style                  `json:"style"`
}

func ParseConfigFile(file string) (*Config, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	conf := &Config{}
	dec := json.NewDecoder(f)
	if err := dec.Decode(conf); err != nil {
		return nil, err
	}
	return conf, nil
}
