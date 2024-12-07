package db

import (
	"database/sql"
	"errors"
	"gardenometer/models"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
  METRIC_TABLE = "metrics"
)

func ReadMetric(db *sql.DB, id uuid.UUID) (*models.Metric, error) {
  result := &models.Metric{}
  row := db.QueryRow("SELECT * FROM " + METRIC_TABLE + " WHERE id = $1", id)
  if (row == nil) {
    return result, errors.New("metric not found")
  }
  err := row.Scan(&result.Id, &result.Name, &result.Moisture, &result.Temp, &result.Lux, &result.UpdatedAt)
  if err == sql.ErrNoRows {
    return nil, nil
  }
  if err != nil {
    return nil, err
  }
  return result, nil
}

func ReadMetricBetweenTimes(db *sql.DB, begin time.Time, end time.Time) ([]*models.Metric, error) {
  rows, err := db.Query("SELECT * FROM " + METRIC_TABLE + " WHERE updated_at BETWEEN $1 AND $2 ORDER BY updated_at ASC", begin, end)
  if err != nil {
    return nil, err
  }
  defer rows.Close()
  result := make([]*models.Metric, 0);
  for rows.Next() {
    m := &models.Metric{}
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

func ReadLatestMetricForEachName(db *sql.DB) ([]*models.RegistrationList, error) {
  sb := strings.Builder{}
  sb.WriteString("SELECT id, reg.name as name, is_active, moisture, temp, lux, ")
  sb.WriteString("reg.updated_at as reg_ts, met.updated_at as met_ts FROM (")
  sb.WriteString("SELECT DISTINCT ON (\"name\") * FROM ")
  sb.WriteString(METRIC_TABLE)
  sb.WriteString(" ORDER BY name, updated_at DESC) as met JOIN ")
  sb.WriteString(REGISTRATION_TABLE)
  sb.WriteString(" as reg ON reg.name = met.name")
  rows, err := db.Query(sb.String())
  if err != nil {
    return nil, err
  }
  defer rows.Close()
  result := make([]*models.RegistrationList, 0)
  for rows.Next() {
    m := &models.RegistrationList{}
    err = rows.Scan(&m.Id, &m.Name,
      &m.IsActive, &m.Moisture, &m.Temp, &m.Lux,
      &m.RegistrationUpdatedAt, &m.UpdatedAt)
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
  defer stmt.Close()
  return err
}
