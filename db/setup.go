package db

import (
  "database/sql"
  _ "github.com/lib/pq"
)

// Setup the database connection.
func Setup() (*sql.DB, error) {
  // TODO move this info to config file
  connString := "user=gardenometer_user dbname=gardenometer_db password=password sslmode=disable"
  return sql.Open("postgres", connString)
}
