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
	DBHost       string
	DBPasswd     string
}

var config Configuration

func loadConfig() {
	configfile := "config.json"
	switch os.Getenv("SMECONFIG") {
	case "production":
		configfile = "config-prod.json"
	}
	file, err := os.Open(configfile)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	if len(config.DBPasswd) == 0 {
		config.DBPasswd = os.Getenv("POSTGRES_PASSWORD")
	}
}
