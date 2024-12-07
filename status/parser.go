package status

import (
	"gardenometer/models"
	"strconv"
	"strings"
)

type Payload struct {
  data []byte
}

func (sp *Payload) Write(p []byte) (n int, err error) {
  if (sp.data == nil) {
    sp.data = make([]byte, 0)
  }
  sp.data = append(sp.data, p...)
  return len(sp.data), nil
}

func (sp *Payload) ParseStatus() models.Status {
  status := string(sp.data)
  result := models.Status{}
  status, _ = strings.CutPrefix(status, "status:")
  values := strings.Split(status, ";")
  for _, entry := range values {
    components := strings.Split(entry, "=")
    if len(components) == 2 {
      prop := components[0]
      val := components[1]
      switch (prop) {
      case "l": {
        l, err := strconv.ParseFloat(val, 64)
        if (err == nil) {
          result.Lux = new(float64)
          *result.Lux = l
        }
        break;
      }
      case "m": {
        l, err := strconv.Atoi(val)
        if (err == nil) {
          result.Moisture = new(int)
          *result.Moisture = l
        }
        break;
      }
      case "t": {
        l, err := strconv.ParseFloat(val, 64)
        if (err == nil) {
          result.Temp = new(float64)
          *result.Temp = l
        }
        break;
      }
      case "e": {
        result.Err = new(string)
        *result.Err = val
        break;
      }
      case "id": {
        result.Id = strings.TrimSpace(val)
        break;
      }
      }
    }
  }
  return result
}
