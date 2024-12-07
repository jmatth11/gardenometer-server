package routes

import (
	"database/sql"
	"fmt"
	"gardenometer/db"
	"gardenometer/models"
	"gardenometer/status"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func home(c echo.Context) error {
  fmt.Println("calling from: ", c.Request().Header.Get("User-Agent"))
  return c.String(http.StatusOK, "gardenometer")
}

func ping(c echo.Context) error {
  fmt.Println("calling from: ", c.Request().Header.Get("User-Agent"))
  return c.String(http.StatusOK, "gardenometer")
}

func createStatus(conn *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    body := c.Request().Body
    sp := new(status.Payload)
    _, err := io.Copy(sp, body)
    if err != nil {
      log.Println(err)
      return c.String(http.StatusInternalServerError, "error handling request")
    }
    status := sp.ParseStatus()
    if (status.Err != nil) {
      log.Println(err)
    } else {
      err = status.ValidateForInsert()
      if err != nil {
        log.Println(err)
        return c.String(http.StatusBadRequest, "must include all metric values")
      }
      metric := models.Metric{}
      metric.FromStatus(status)
      if err := db.InsertMetric(conn, metric); err != nil {
        c.String(http.StatusInternalServerError, err.Error())
      }
    }
    return c.String(http.StatusOK, "")
  }
}

func createQueue() echo.HandlerFunc {
  return func(c echo.Context) error {
    return nil
  }
}

func Setup(e *echo.Echo, conn *sql.DB) {
  e.GET("/", home)
  e.GET("/status", ping)
  e.POST("/status", createStatus(conn))
  e.POST("/queue", createQueue())
}
