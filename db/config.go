package db

import (
	"database/sql"
	"errors"
	"gardenometer/models"
)

const (
  CONFIG_TABLE = "config"
)

func ReadConfigForDevice(conn *sql.DB, name string) (*models.ConfigData, error) {
  row := conn.QueryRow("SELECT * FROM " + CONFIG_TABLE + " WHERE name = $1", name)
  if row == nil {
    return nil, errors.New("config not found")
  }
  result := &models.ConfigData{}
  err := row.Scan(&result.Name, &result.MoistureAir, &result.MoistureWater, &result.Wait)
  if err == sql.ErrNoRows {
    return nil, nil
  }
  if err != nil {
    return nil, err
  }
  return result, nil
}

func InsertConfigForDevice(conn *sql.DB, obj *models.ConfigData) error {
  stmt, err :=  conn.Prepare("INSERT INTO " + CONFIG_TABLE +
    " (name, moisture_air, moisture_water, wait_time) VALUES ($1, $2, $3, $4)")
  if err != nil {
    return err
  }
  _, err = stmt.Exec(obj.Name, obj.MoistureAir, obj.MoistureWater, obj.Wait)
  defer stmt.Close()
  return err
}

func UpdateConfigForDevice(conn *sql.DB, obj *models.ConfigData) error {
  stmt, err :=  conn.Prepare("UPDATE " + CONFIG_TABLE +
    " SET moisture_air=$2, moisture_water=$3, wait_time=$4 WHERE name=$1")
  if err != nil {
    return err
  }
  _, err = stmt.Exec(obj.Name, obj.MoistureAir, obj.MoistureWater, obj.Wait)
  defer stmt.Close()
  return err
}
