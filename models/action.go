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
  Values map[string]interface{}
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
    sb.WriteString(fmt.Sprintf("%s=%s", key, value))
  }
  return sb.String()
}

