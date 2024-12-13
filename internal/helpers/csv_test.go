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
