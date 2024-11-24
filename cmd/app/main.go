package main

import (
	"fmt"
	"log"
	"onse/internal/helpers"
)

func main() {
	// load app configuration
	config := helpers.NewConfig()
	if config.DebugMode {
		fmt.Printf("config is %+v\n", config)
	}

	// read and parse the CSV file
	participants, err := helpers.GetParticipants(config.CsvFile)
	if err != nil {
		log.Fatalf("error reading CSV: %s", err)
	}

	for _, p := range participants {
		fmt.Printf("Participant %d has email %s\n", p.Id, p.Email)
	}
}
