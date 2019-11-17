package domain

import (
	"time"

	"github.com/mitchellh/mapstructure"
)

type Event struct {
	ID          string      `json:"event_id"`
	EventType   string      `json:"event_type"`
	AggregateID string      `json:"aggregate_id"`
	Payload     interface{} `json:"payload"`
	OccurredOn  time.Time   `json:"occurred_on"`
}

func (e *Event) DecodePayload(i interface{}) error {
	return mapstructure.Decode(e.Payload, i)
}
