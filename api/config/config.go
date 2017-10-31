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

// App TODO
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

	// Overwrite config
	if len(os.Getenv("MIMIR_DB_HOST")) > 0 {
		App.DB.Host = os.Getenv("MIMIR_DB_HOST")
		App.DB.Username = os.Getenv("MIMIR_DB_USERNAME")
		App.DB.Password = os.Getenv("MIMIR_DB_PASSWORD")
		if os.Getenv("MIMIR_DB_SSL") == "true" {
			App.DB.SSL = true
		}
	}
}
