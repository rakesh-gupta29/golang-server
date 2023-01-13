package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	Address    string `env:"ADDRESS" envDefault:"localhost:4000"`
	FromEmail  string `env:"EMAIL"`
	Password   string `env:"PASSWORD"`
	SMTPServer string `env:"SERVER"`
	SMTPPort   string `env:"PORT"`
	Target     string `env:"TARGET"`
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load("config.env") // load the  file path from local file.
	// Specify path file if .env file does not reside in same directory as main.go
	if err != nil {
		log.Println("Error Loading ENV Variables.", err)
	}
	cfg := Config{}
	err = env.Parse(&cfg) // parsing to the struct
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
