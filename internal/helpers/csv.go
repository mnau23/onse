package helpers

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	mail_checker "github.com/FGRibreau/mailchecker/v6/platform/go"
)

type Participant struct {
	Id         int
	Name       string
	Email      string
	Exclusions []int
	Message    string
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
		record, err := reader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("failed to read row: %w", err)
		}

		// parse row elements
		participantId, _ := strconv.Atoi(record[0])
		exclusions, _ := getExclusions(record[3])

		// validate email address syntax and if disposable
		email := record[2]
		if !mail_checker.IsValid(email) {
			return nil, fmt.Errorf("found invalid participant email '%s'", email)
		}

		participants = append(participants, Participant{
			Id:         participantId,
			Name:       record[1],
			Email:      email,
			Exclusions: exclusions,
			Message:    record[4],
		})
	}

	return participants, nil
}

// converts the strings with a JSON array into a slice
func getExclusions(field string) ([]int, error) {
	var exclusions []int

	if err := json.Unmarshal([]byte(field), &exclusions); err != nil {
		return nil, fmt.Errorf("failed to parse exclusions: %w", err)
	}

	return exclusions, nil
}
