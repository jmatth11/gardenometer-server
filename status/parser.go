package status

import (
	"gardenometer/models"
	"strconv"
	"strings"
)

func ParseStatus(status string) models.Status {
  result := models.Status{}
  values := strings.Split(status, ";")
  for _, entry := range values {
    components := strings.Split(entry, "=")
    if len(components) == 2 {
      prop := components[0]
      val := components[1]
      switch (prop) {
      case "l": {
        l, err := strconv.ParseFloat(val, 64)
        if (err != nil) {
          result.Lux = &l
        }
        break;
      }
      case "m": {
        l, err := strconv.Atoi(val)
        if (err != nil) {
          result.Moisture = &l
        }
        break;
      }
      case "t": {
        l, err := strconv.ParseFloat(val, 64)
        if (err != nil) {
          result.Temp = &l
        }
        break;
      }
      case "e": {
        result.Err = &val
        break;
      }
      case "id": {
        result.Id = val
        break;
      }
      }
    }
  }
  return result
}
