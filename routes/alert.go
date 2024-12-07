package routes

import (
	"database/sql"
	"gardenometer/db"
	"gardenometer/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func createAlert(conn *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    payload := new(models.Alerts)
    if err := c.Bind(payload); err != nil {
      return c.String(http.StatusBadRequest, "invalid payload")
    }
    if err := models.ValidateAlert(payload); err != nil {
      log.Println(err)
      return c.String(http.StatusBadRequest, "invalid payload")
    }
    reg, err := db.ReadRegistration(conn, payload.Name)
    if err != nil {
      log.Println(err)
      return c.String(http.StatusInternalServerError, "registration read error")
    }
    if reg == nil || reg.Name == "" {
      log.Println("not a valid registration name")
      return c.String(http.StatusBadRequest, "not a valid registration name")
    }
    err = db.InsertAlert(conn, payload)
    if err != nil {
      log.Println(err)
      return c.String(http.StatusInternalServerError, "alert insert error")
    }
    return c.String(http.StatusOK, "")
  }
}

