package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetParticipants(t *testing.T) {
	participants, err := GetParticipants("./testdata/participants.csv")
	assert.NoError(t, err)

	assert.Len(t, participants, 4)

	assert.Equal(t, 1, participants[0].Id)
	assert.Equal(t, "John", participants[1].Name)
	assert.Equal(t, "alice@fake.com", participants[2].Email)
	assert.Equal(t, []int{2, 3}, participants[3].Exclusions)
	assert.Equal(t, "whatsup", participants[3].Message)
}

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"test@example.com", true},  // valid
		{"invalid-email", false},    // invalid
		{"@wrong.com", false},       // invalid
		{"user@domain..com", false}, // invalid
	}

	for _, test := range tests {
		t.Run(test.email, func(t *testing.T) {
			result := IsValidEmail(test.email)
			if result != test.expected {
				t.Errorf("IsValidEmail(%q): %v - expected: %v", test.email, result, test.expected)
			}
		})
	}
}
