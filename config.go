package main

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var config Configuration

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
}
