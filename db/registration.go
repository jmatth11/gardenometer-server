package db

import (
	"database/sql"
	"errors"
	"gardenometer/models"
)

const (
  REGISTRATION_TABLE = "registrations"
)

func ReadRegistration(db *sql.DB, name string) (*models.Registration, error) {
  result := &models.Registration{}
  row := db.QueryRow("SELECT * FROM " + REGISTRATION_TABLE + " WHERE name = $1", name)
  if (row == nil) {
    return result, errors.New("registration not found")
  }
  err := row.Scan(&result.Name, &result.IsActive, &result.UpdatedAt)
  if err == sql.ErrNoRows {
    return nil, nil
  }
  if err != nil {
    return nil, err
  }
  return result, nil
}

func ReadAllRegistration(db *sql.DB) ([]models.Registration, error) {
  rows, err := db.Query("SELECT * FROM " + REGISTRATION_TABLE)
  if err != nil {
    return nil, err
  }
  defer rows.Close()
  result := make([]models.Registration, 0)
  for rows.Next() {
    tmp := models.Registration{}
    err = rows.Scan(&tmp.Name, &tmp.IsActive, &tmp.UpdatedAt)
    if err != nil {
      return nil, err
    }
    result = append(result, tmp)
  }
  err = rows.Err()
  if err != nil {
    return nil, err
  }
  return result, nil
}

func InsertRegistration(db *sql.DB, reg *models.Registration) error {
  stmt, err := db.Prepare("INSERT INTO " + REGISTRATION_TABLE + " (name, is_active, updated_at) VALUES ( $1, $2, $3 )")
  if err != nil {
    return err
  }
  defer stmt.Close()
  _, err = stmt.Exec(reg.Name, reg.IsActive, reg.UpdatedAt)
  return err
}

func UpdateRegistrationIsActive(conn *sql.DB, name string, isActive bool) error {
  stmt, err := conn.Prepare("UPDATE " + REGISTRATION_TABLE + " SET is_active = $1 WHERE name = $2")
  if err != nil {
    return err
  }
  defer stmt.Close()
  _, err = stmt.Exec(isActive, name)
  return err
}
