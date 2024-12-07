package db

import (
	"time"

	"github.com/google/uuid"
)

type Metric struct {
  Id uuid.UUID
  Name string
  Moisture int
  Temp float64
  Lux float64
  UpdatedAt time.Time
}

func ReadMetric(id uuid.UUID) (Metric, error) {
  // TODO flesh out db calls
}
