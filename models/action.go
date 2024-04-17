package models

import (
	"fmt"
	"strings"
)

type ActionType string

const (
  ActionCalibrate ActionType = "cal"
  ActionCode ActionType = "code"
  ActionConfig ActionType = "config"
)

type Action struct {
  Type ActionType
  Name string
  Values map[string]interface{}
}

func NewAction(t ActionType, name string) Action {
  return Action{
    Type: t,
    Name: name,
    Values: make(map[string]interface{}),
  }
}

func (a *Action) String() string {
  sb := strings.Builder{}
  sb.WriteString(string(a.Type))
  sb.WriteString(":")
  first := true
  for key, value := range a.Values {
    if !first {
      sb.WriteString(";")
    }
    if a.Type == ActionCode {
      sb.WriteString(key)
    } else {
      sb.WriteString(fmt.Sprintf("%s=%s", key, value))
    }
  }
  return sb.String()
}

