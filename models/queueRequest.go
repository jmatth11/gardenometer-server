package models

type QueueRequest struct {
  Type ActionType `json:"type"`
  Name string `json:"name"`
  Configuration *Config `json:"configuration"`
  Code *int `json:"code"`
}
