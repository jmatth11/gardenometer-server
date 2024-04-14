package models

import (
	"errors"
	"strings"
)

// Codes values to send to arduino
const (
  CodeNone = iota
  CodeError
  CodeClearError
)

// Status object sent from arduino
type Status struct {
  Id string
  Temp *float64
  Lux *float64
  Moisture *int
  Err *string
}

func (s *Status) ValidateForInsert() error {
  sb := strings.Builder{}
  if s.Id == "" {
    sb.WriteString("Id was empty; ")
  }
  if s.Temp == nil {
    sb.WriteString("temp was nil; ")
  }
  if s.Lux == nil {
    sb.WriteString("lux was nil; ")
  }
  if s.Moisture == nil {
    sb.WriteString("moisture was nil; ")
  }
  if (sb.Len() > 0) {
    return errors.New(sb.String())
  }
  return nil
}

