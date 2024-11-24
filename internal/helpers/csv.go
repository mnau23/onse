package helpers

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// reads a CSV file and returns a slice of Participant objects
func GetParticipants(filename string) ([]Participant, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// read and discard the header row
	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read header row: %w", err)
	}

	var participants []Participant
	for {
		row, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("failed to read row: %w", err)
		}

		// parse row elements
		participant_id, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse ID: %w", err)
		}
		participant_email := row[1]

		participants = append(participants, Participant{Id: participant_id, Email: participant_email})
	}

	return participants, nil
}
