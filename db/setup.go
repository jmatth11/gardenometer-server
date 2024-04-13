package db

import "database/sql"

// Setup the database connection.
func Setup() (*sql.DB, error) {
  return sql.Open("sqlite3", "./db.sqlite")
}
