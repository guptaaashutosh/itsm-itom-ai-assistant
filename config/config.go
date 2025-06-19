// Package config loads application configuration using Viper from the .env file.
package config

import (
	"log"

	"github.com/spf13/viper"
)

var Conf Config

// Config holds all application configuration loaded from environment variables.
type Config struct {
	Port        string `mapstructure:"PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

// LoadConfig loads environment variables from .config/.env into a Config struct.
func LoadConfig() (Config, error) {
	viper.SetConfigFile(".config/.env")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("no config file found. going on...")
		} else {
			log.Printf("error loading config file")
		}
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		return Config{}, err
	}

	return Conf, err
}
