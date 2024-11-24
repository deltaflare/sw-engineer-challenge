package service_test

import (
	"testing"
	"time"

	"github.com/OtchereDev/deltaflare/daemon/service"
	"github.com/stretchr/testify/assert"
)

func TestCreateEvent(t *testing.T) {
	event := service.CreateEvent()

	assert.GreaterOrEqual(t, event.Criticality, 1)
	assert.LessOrEqual(t, event.Criticality, 5)

	assert.Equal(t, event.EventMessage, "New event created")

	assert.NotEqual(t, event.Timestamp, time.Time{})
}
