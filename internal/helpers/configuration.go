package helpers

import (
	"log"
	"onse/internal/email"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	DebugMode     bool   `env:"DEBUG_MODE" envDefault:"false"`
	CsvFile       string `env:"CSV_FILE,required"`
	Smtp          email.SmtpConfig
	EmailTemplate string `env:"EMAIL_TEMPLATE,required"`
	EmailSender   string `env:"EMAIL_SENDER,required"`
}

func NewConfig() Config {
	// load the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("unable to load .env file: %v", err)
	}

	// parse vars into Config
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("unable to parse env vars: %v", err)
	}

	return cfg
}
