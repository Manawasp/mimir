package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Config TODO
type Config struct {
	DB    Database  `toml:"database"`
	Hooks []Webhook `toml:"webhook"`
}

var App Config

// Init conf var read toml file and add it inside of the var
func init() {
	configFile := "config.toml"
	_, err := os.Stat(configFile)
	if err != nil {
		log.Fatal("Config file is missing.")
	}

	if _, err := toml.DecodeFile(configFile, &App); err != nil {
		log.Fatal(err)
	}
}
