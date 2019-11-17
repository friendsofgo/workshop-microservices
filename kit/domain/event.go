package domain

import (
	"encoding/json"
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

func (e *Event) decodePayload(i interface{}) error {
	return mapstructure.Decode(e.Payload, i)
}

func EventDecode(message []byte, payload interface{}) (Event, error) {
	var decoded Event
	err := json.Unmarshal(message, &decoded)
	if err != nil {
		return Event{}, nil
	}

	if err := decoded.decodePayload(payload); err != nil {
		return Event{}, nil
	}
	return decoded, nil
}
