package services_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/OtchereDev/deltaflare/cient/models"
	"github.com/OtchereDev/deltaflare/cient/services"
	"github.com/stretchr/testify/assert"
)

// TestDisplayEvent tests the DisplayEvent function
func TestDisplayEvent(t *testing.T) {
	event := models.Event{
		Criticality:  5,
		EventMessage: "Test event",
		Timestamp:    time.Now().Format(time.RFC3339),
	}

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}

	old := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = old }()

	services.DisplayEvent(event)

	w.Close()

	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("failed to read from pipe: %v", err)
	}

	expectedOutput := fmt.Sprintf("Criticality: %d | EventMessage: %s | Timestamp: %s\n",
		event.Criticality, event.EventMessage, event.Timestamp)

	assert.Equal(t, expectedOutput, buf.String())
}
