package actions

import (
	"fmt"
	"gardenometer/models"
)

func GenerateCalibrationAction(name string) models.Action {
  return models.NewAction(models.ActionCalibrate, name)
}

func GenerateConfigAction(name string, c models.Config) models.Action {
  res := models.NewAction(models.ActionConfig, name)
  if c.Wait != nil {
    res.Values[fmt.Sprint(models.ConfigWait)] = *c.Wait
  }
  if c.MoistureAir != nil {
    res.Values[fmt.Sprint(models.ConfigMoistureAir)] = *c.MoistureAir
  }
  if c.MoistureWater != nil {
    res.Values[fmt.Sprint(models.ConfigMoistureWater)] = *c.MoistureWater
  }
  return res
}

func GenerateCodeResponse(name string, code int) models.Action {
  res := models.NewAction(models.ActionCode, name)
  res.Values[fmt.Sprint(code)] = nil
  return res
}
