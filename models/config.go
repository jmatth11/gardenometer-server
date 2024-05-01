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

type ConfigRequest struct {

}

type Config struct {
  Wait *int `json:"wait"`
  Moisture *int `json:"moisture"`
  Temp *int `json:"temp"`
  Lux *int `json:"lux"`
  Cal *int `json:"cal"`
  Err *int `json:"err"`
  Good *int `json:"good"`
}

type ConfigTab struct {
  Devices []string
}
