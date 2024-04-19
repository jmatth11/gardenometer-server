package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Setup the database connection.
func Setup(user string, dbName string, dbPassword string) (*sql.DB, error) {
  connString := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbName, dbPassword)
  return sql.Open("postgres", connString)
}
