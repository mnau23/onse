package helpers

import (
	"log"
	"onse/internal/email"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Smtp email.SmtpConfig
	// data
	CsvFile       string `env:"CSV_FILE,required"`
	EmailTemplate string `env:"EMAIL_TEMPLATE,required"`
	Budget        int    `env:"BUDGET"`
	// testing
	DebugMode         bool   `env:"DEBUG_MODE" envDefault:"false"`
	EmailReceiverTest string `env:"EMAIL_RECEIVER_TEST,required"`
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
