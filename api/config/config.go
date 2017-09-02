package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Config TODO
type Config struct {
	DB   Database  `toml:"database"`
	Hook []Webhook `toml:"webhook"`
}

var AppConfig Config

// Init conf var read toml file and add it inside of the var
func init() {
	configFile := "../config.toml"
	_, err := os.Stat(configFile)
	if err != nil {
		log.Fatal("Config file is missing.")
	}

	if _, err := toml.DecodeFile(configFile, &AppConfig); err != nil {
		log.Fatal(err)
	}
}
