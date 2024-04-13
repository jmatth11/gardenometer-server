package models

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

