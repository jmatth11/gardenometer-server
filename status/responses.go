package status

import (
	"fmt"
	"gardenometer/models"
	"strings"
)

const (
  Calibration = "cal:"
  Config = "config:"
  Code = "code:"
)

func GenerateCalibrationResponse() string {
  return Calibration
}

func GenerateConfigResponse(c models.Config) string {
  sb := strings.Builder{}
  sb.WriteString(Config)
  second := false
  if c.Wait != nil {
    second = true
    sb.WriteString(fmt.Sprintf("%d=%d", models.ConfigWait, *c.Wait))
  }
  if c.Moisture != nil {
    if second {
      sb.WriteString(";")
    }
    second = true
    sb.WriteString(fmt.Sprintf("%d=%d", models.ConfigMoisture, *c.Moisture))
  }
  if c.Lux != nil {
    if second {
      sb.WriteString(";")
    }
    second = true
    sb.WriteString(fmt.Sprintf("%d=%d", models.ConfigLux, *c.Lux))
  }
  if c.Temp != nil {
    if second {
      sb.WriteString(";")
    }
    second = true
    sb.WriteString(fmt.Sprintf("%d=%d", models.ConfigTemp, *c.Temp))
  }
  if c.Cal != nil {
    if second {
      sb.WriteString(";")
    }
    second = true
    sb.WriteString(fmt.Sprintf("%d=%d", models.ConfigCal, *c.Cal))
  }
  if c.Err != nil {
    if second {
      sb.WriteString(";")
    }
    second = true
    sb.WriteString(fmt.Sprintf("%d=%d", models.ConfigErr, *c.Err))
  }
  if c.Good != nil {
    if second {
      sb.WriteString(";")
    }
    second = true
    sb.WriteString(fmt.Sprintf("%d=%d", models.ConfigGood, *c.Good))
  }
  return sb.String()
}

func GenerateCodeResponse(code int) string {
  return fmt.Sprintf("%s%d", Code, code)
}
