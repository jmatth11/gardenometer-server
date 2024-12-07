package models

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type AlertType int

const (
  AlertMoisture AlertType = iota
  AlertTemp
)

func (a AlertType) String() string {
  switch (a) {
    case AlertMoisture:
      return "Moisture"
    case AlertTemp:
      return "Temperature"
  }
  return ""
}

type Alerts struct {
  Id uuid.UUID `json:"id"`
  Name string `json:"name"`
  Key AlertType `json:"key"`
  Value float64 `json:"value"`
}

func ValidateAlert(a *Alerts) error {
  sb := strings.Builder{}
  if a.Name == "" {
    sb.WriteString("registration name cannot be empty")
  }
  if a.Key < AlertMoisture || a.Key > AlertTemp {
    sb.WriteString("invalid alert key")
  }
  if sb.Len() > 0 {
    return errors.New(sb.String())
  }
  return nil
}
