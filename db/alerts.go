package db

import (
	"database/sql"
	"gardenometer/models"

	"github.com/google/uuid"
)

const (
  ALERTS_TABLE = "alerts"
)

func ReadAlertsForName(conn *sql.DB, name string) ([]*models.Alerts, error) {
  rows, err := conn.Query("SELECT * FROM " + ALERTS_TABLE + " WHERE name = $1", name)
  if err != nil {
    return nil, err
  }
  defer rows.Close()
  result := make([]*models.Alerts, 0);
  for rows.Next() {
    m := &models.Alerts{}
    err = rows.Scan(&m.Id, &m.Name, &m.Key, &m.Value)
    if err != nil {
      return nil, err
    }
    result = append(result, m)
  }
  err = rows.Err()
  if err != nil {
    return nil, err
  }
  return result, nil
}

func InsertAlert(conn *sql.DB, obj *models.Alerts) error {
  obj.Id = uuid.New()
  stmt, err := conn.Prepare("INSERT INTO " + ALERTS_TABLE +
  " (id, name, key_name, value) VALUES ($1, $2, $3, $4)")
  if err != nil {
    return err
  }
  _, err = stmt.Exec(obj.Id, obj.Name, obj.Key, obj.Value)
  defer stmt.Close()
  return err
}
