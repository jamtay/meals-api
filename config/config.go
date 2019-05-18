package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Represents database server and credentials
type Config struct {
	Server   string
	Database string
}

// ReadMoviesConfig and parse the configuration file
func (c *Config) ReadMealsConfig() {
	if _, err := toml.DecodeFile("config/meals_config.toml", &c); err != nil {
		log.Fatal(err)
	}
}