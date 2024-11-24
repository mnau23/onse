package main

import (
	"fmt"
	"log"
	"onse/internal/helpers"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	CsvFile string `env:"CSV_FILE,required"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	cfg := Config{}
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("unable to parse env vars: %e", err)
	}

	// read and parse the CSV file
	participants, err := helpers.GetParticipants(cfg.CsvFile)
	if err != nil {
		log.Fatalf("error reading CSV: %s", err)
	}

	for _, p := range participants {
		fmt.Printf("Participant %d has email %s\n", p.Id, p.Email)
	}
}
