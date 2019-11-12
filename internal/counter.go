package counter

import (
	"time"
)

type Counter struct {
	ID         string
	Name       string
	Value      uint
	BelongsTo  string
	UpdatedAt *time.Time
}



