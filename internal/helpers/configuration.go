package helpers

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	CsvFile string `env:"CSV_FILE,required"`
}

func NewConfig() Config {
	// load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %v", err)
	}

	// parse vars into Config
	cfg := Config{}
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("unable to parse env vars: %v", err)
	}

	return cfg
}
