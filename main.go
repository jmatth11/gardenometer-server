package main

import (
	"gardenometer/db"
	"gardenometer/routes"

	"github.com/labstack/echo/v4"
)

func main() {
  e := echo.New()
  db, err := db.Setup()
  if err != nil {
    panic(err)
  }
  routes.Setup(e, db);
  e.Logger.Fatal(e.Start(":8000"))
}
