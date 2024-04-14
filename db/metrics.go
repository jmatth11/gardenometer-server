package db

import (
	"database/sql"
	"errors"
	"gardenometer/models"
	"time"

	"github.com/google/uuid"
)

const (
  METRIC_TABLE = "metrics"
)

func ReadMetric(db *sql.DB, id uuid.UUID) (models.Metric, error) {
  result := models.Metric{}
  row := db.QueryRow("SELECT * FROM " + METRIC_TABLE + " WHERE id = $1", id)
  if (row == nil) {
    return result, errors.New("metric not found")
  }
  row.Scan(&result.Id, &result.Name, &result.Moisture, &result.Temp, &result.Lux, &result.UpdatedAt)
  return result, nil
}

func ReadMetricBetweenTimes(db *sql.DB, begin time.Time, end time.Time) ([]models.Metric, error) {
  rows, err := db.Query("SELECT * FROM " + METRIC_TABLE + " WHERE updated_at BETWEEN $1 AND $2", begin, end)
  if err != nil {
    return nil, err
  }
  defer rows.Close()
  result := make([]models.Metric, 0);
  for rows.Next() {
    m := models.Metric{}
    err = rows.Scan(&m.Id, &m.Name, &m.Moisture, &m.Temp, &m.Lux, &m.UpdatedAt)
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

func InsertMetric(db *sql.DB, m models.Metric) error {
  m.Id = uuid.New()
  m.UpdatedAt = time.Now()
  stmt, err := db.Prepare("INSERT INTO " + METRIC_TABLE + " (id, name, moisture, temp, lux, updated_at) VALUES ($1, $2, $3, $4, $5, $6)")
  if err != nil {
    return err
  }
  _, err = stmt.Exec(m.Id, m.Name, m.Moisture, m.Temp, m.Lux, m.UpdatedAt)
  if err != nil {
    return err
  }
  defer stmt.Close()
  return nil
}
