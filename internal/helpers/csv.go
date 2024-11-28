package helpers

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Participant struct {
	Id      int
	Email   string
	Name    string
	Message string
}

// reads a CSV file and returns a slice of Participant objects
func GetParticipants(filename string) ([]Participant, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// read and discard the header row
	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("failed to read header row: %w", err)
	}

	var participants []Participant
	for {
		row, err := reader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("failed to read row: %w", err)
		}

		// parse row elements
		participantId, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse ID: %w", err)
		}

		participants = append(participants, Participant{
			Id:      participantId,
			Email:   row[1],
			Name:    row[2],
			Message: row[3],
		})
	}

	return participants, nil
}
