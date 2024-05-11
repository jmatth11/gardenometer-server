package routes

import (
	"database/sql"
	"gardenometer/db"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func getDisplay(conn *sql.DB) echo.HandlerFunc {
  return func (c echo.Context) error {
    id := c.Param("id")
    if id == "" {
      return c.Render(http.StatusBadRequest, "error", "Must supply an ID")
    }
    timeStr := c.Param("time")
    if timeStr == "" {
      return c.Render(http.StatusBadRequest, "error", "Must supply a start time")
    }
    start, err := time.Parse(time.DateOnly, timeStr)
    if err != nil {
      log.Println(err)
      return c.Render(http.StatusBadRequest, "toast", ErrorToast(err.Error()))
    }
    req, err := db.ReadRegistration(conn, id)
    if err != nil {
      log.Println(err)
      return c.Render(http.StatusInternalServerError, "error", err)
    }
    if req == nil {
      return c.Render(
        http.StatusNotFound,
        "toast",
        ErrorToast("Device is not registered"),
      )
    }
    metrics, err := db.ReadMetricBetweenTimes(conn, start, time.Now())
    if err != nil {
      log.Println(err)
      return c.Render(http.StatusInternalServerError, "error", err)
    }
    return c.JSON(http.StatusOK, metrics)
  }
}
