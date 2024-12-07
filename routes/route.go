package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ping(c echo.Context) error {
  fmt.Println("calling from: ", c.Request().Header.Get("User-Agent"))
  return c.String(http.StatusOK, "gardenometer")
}

func createStatus() echo.HandlerFunc {
  return func(c echo.Context) error {
    return nil
  }
}

func Setup(e *echo.Echo, db *sql.DB) {
  e.GET("/status", ping)
  e.POST("/status", createStatus())
}
