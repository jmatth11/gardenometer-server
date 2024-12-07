package models

import (
	"strconv"
	"strings"
)

const (
  ConfigWait = iota
  ConfigMoistureAir
  ConfigMoistureWater
)

type ConfigRequest struct {
  data []byte
}

type Config struct {
  Name string `json:"name"`
  Wait *int `json:"wait"`
  MoistureAir *int `json:"moisture_air"`
  MoistureWater *int `json:"moisture_water"`
}

type ConfigData struct {
  Name string `json:"name"`
  Wait int `json:"wait"`
  MoistureAir int `json:"moisture_air"`
  MoistureWater int `json:"moisture_water"`
}

type ConfigTab struct {
  Devices []string
  Configurations map[string]*ConfigData
}

func (c *ConfigRequest) Write(p []byte) (n int, err error) {
  if (c.data == nil) {
    c.data = make([]byte, 0)
  }
  c.data = append(c.data, p...)
  return len(c.data), nil
}

func (c *ConfigRequest) ParseConfig() *Config {
  configStr := string(c.data)
  result := &Config{}
  configStr, _ = strings.CutPrefix(configStr, "config:")
  values := strings.Split(configStr, ";")
  for _, entry := range values {
    components := strings.Split(entry, "=")
    if len(components) == 2 {
      prop := components[0]
      val := components[1]
      switch (prop) {
        case "id": {
          result.Name = strings.TrimSpace(val)
          break;
        }
        case "mw": {
          l, err := strconv.Atoi(val)
          if (err == nil) {
            result.MoistureWater = new(int)
            *result.MoistureWater = l
          }
          break;
        }
        case "ma": {
          l, err := strconv.Atoi(val)
          if (err == nil) {
            result.MoistureAir = new(int)
            *result.MoistureAir = l
          }
          break;
        }
        case "wt": {
          l, err := strconv.Atoi(val)
          if (err == nil) {
            result.Wait = new(int)
            *result.Wait = l
          }
          break;
        }
      }
    }
  }
  return result
}

func (cd *ConfigData) ToConfig() Config {
  result := Config{
    Wait: new(int),
    MoistureAir: new(int),
    MoistureWater: new(int),
  }
  *result.Wait = cd.Wait
  *result.MoistureAir = cd.MoistureAir
  *result.MoistureWater = cd.MoistureWater
  return result
}

