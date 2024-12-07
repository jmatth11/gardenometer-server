package models

// Configuration enum for addressing which pin to set
const (
  ConfigWait = iota
  ConfigMoisture
  ConfigLux
  ConfigTemp
  ConfigCal
  ConfigErr
  ConfigGood
)

type Config struct {
  Wait *int
  Moisture *int
  Temp *int
  Lux *int
  Cal *int
  Err *int
  Good *int
}
