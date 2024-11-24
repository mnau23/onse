package main

import (
	"fmt"
	"log"

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

	fmt.Printf("CSV: %s\n", cfg.CsvFile)
}
