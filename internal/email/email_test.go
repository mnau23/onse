package email

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHtml(t *testing.T) {
	expected, _ := os.ReadFile("./testdata/expected.html")

	templatePath, _ := filepath.Abs(filepath.Join("./../../data/email_template.html"))
	data := EmailData{
		GifterName:   "John Doe",
		GifterEmail:  "fake@email.com",
		ReceiverName: "Jane Doe",
		Message:      "Test message",
	}
	result, err := ParseHtml(templatePath, data)

	assert.NoError(t, err)
	assert.Equal(t, string(expected), result, "Output should match expected result")
}
