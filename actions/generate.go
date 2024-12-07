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
  if c.Moisture != nil {
    res.Values[fmt.Sprint(models.ConfigMoisture)] = *c.Moisture
  }
  if c.Lux != nil {
    res.Values[fmt.Sprint(models.ConfigLux)] = *c.Lux
  }
  if c.Temp != nil {
    res.Values[fmt.Sprint(models.ConfigTemp)] = *c.Temp
  }
  if c.Cal != nil {
    res.Values[fmt.Sprint(models.ConfigCal)] = *c.Cal
  }
  if c.Err != nil {
    res.Values[fmt.Sprint(models.ConfigErr)] = *c.Err
  }
  if c.Good != nil {
    res.Values[fmt.Sprint(models.ConfigGood)] = *c.Good
  }
  return res
}

func GenerateCodeResponse(name string, code int) models.Action {
  res := models.NewAction(models.ActionCode, name)
  res.Values[fmt.Sprint(code)] = nil
  return res
}
