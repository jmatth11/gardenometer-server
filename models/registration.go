package models

import (
	"errors"
	"strings"
	"time"
)

type Registration struct {
  Name string `json:"name"`
  IsActive bool `json:"isActive"`
  UpdatedAt time.Time `json:"updatedAt"`
}

func NewRegistration(name string) *Registration {
  return &Registration{
    Name: name,
    IsActive: true,
    UpdatedAt: time.Now(),
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
