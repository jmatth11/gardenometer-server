package models

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Registration struct {
  Name string `json:"name"`
  IsActive bool `json:"isActive"`
  UpdatedAt time.Time `json:"updatedAt"`
}

type RegistrationList struct {
  Id uuid.UUID
  IsActive bool
  Name string
  Moisture int
  Temp float64
  Lux float64
  RegistrationUpdatedAt time.Time
  UpdatedAt time.Time
}

func NewRegistration(name string) *Registration {
  return &Registration{
    Name: name,
    IsActive: true,
    UpdatedAt: time.Now().UTC(),
  }
}

func ValidateRegistration(reg *Registration) error {
  sb := strings.Builder{}
  if reg.Name == "" {
    sb.WriteString("name is required")
  }
  if sb.Len() > 0 {
    return errors.New(sb.String())
  }
  return nil
}
