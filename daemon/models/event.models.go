package models

import "time"

type Event struct {
	Criticality  int       `json:"criticality"`
	Timestamp    time.Time `json:"timestamp"`
	EventMessage string    `json:"event_message"`
}
