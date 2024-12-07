package models

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

func (m *Metric) FromStatus(s Status) Metric {
  m.Name = s.Id
  m.Moisture = *s.Moisture
  m.Temp = *s.Temp
  m.Lux = *s.Lux
  return *m
}

