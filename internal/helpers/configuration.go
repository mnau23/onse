package helpers

import (
	"log"
	"onse/internal/email"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	DebugMode   bool   `env:"DEBUG_MODE" envDefault:"false"`
	CsvFile     string `env:"CSV_FILE,required"`
	EmailSender email.EmailSender
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
