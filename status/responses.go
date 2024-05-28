package status

import (
	"fmt"
	"gardenometer/models"
	"strings"
)

func GenerateCalibrationResponse() string {
  return fmt.Sprintf("%s:", string(models.ActionCalibrate))
}

func GenerateConfigResponse(c models.Config) string {
  sb := strings.Builder{}
  sb.WriteString(string(models.ActionConfig))
  sb.WriteString(":")
  second := false
  if c.Wait != nil {
    second = true
    sb.WriteString(fmt.Sprintf("%d=%d", models.ConfigWait, *c.Wait))
  }
  if c.MoistureAir != nil {
    if second {
      sb.WriteString(";")
    }
    second = true
    sb.WriteString(fmt.Sprintf("%d=%d", models.ConfigMoistureAir, *c.MoistureAir))
  }
  if c.MoistureWater != nil {
    if second {
      sb.WriteString(";")
    }
    second = true
    sb.WriteString(fmt.Sprintf("%d=%d", models.ConfigMoistureWater, *c.MoistureWater))
  }
  return sb.String()
}

func GenerateCodeResponse(code int) string {
  return fmt.Sprintf("%s%d", models.ActionCode, code)
}
