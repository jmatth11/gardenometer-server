package routes

import (
	"database/sql"
	"fmt"
	"gardenometer/actions"
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

func createStatus(conn *sql.DB, actionCache actions.Queue) echo.HandlerFunc {
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
      reg, err := db.ReadRegistration(conn, status.Id)
      if err != nil {
        log.Println(err)
        return c.String(http.StatusInternalServerError, "registration read failed")
      }
      if err := models.ValidateRegistration(reg); err != nil {
        newReq := models.NewRegistration(status.Id)
        err = db.InsertRegistration(conn, newReq)
        if err != nil {
          log.Println(err)
          return c.String(http.StatusInternalServerError, "registration insertion failed")
        }
      }
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
    a, found := actionCache.NextActionForName(status.Id)
    result := ""
    if found {
      result = a.String()
    }
    return c.String(http.StatusOK, result)
  }
}

func createQueue(actionCache actions.Queue) echo.HandlerFunc {
  return func(c echo.Context) error {
    qr := new(models.QueueRequest)
    if err := c.Bind(qr); err != nil {
      return c.String(http.StatusBadRequest, "invalid payload")
    }
    switch(qr.Type) {
      case models.ActionCalibrate:{
        actionCache.Push(actions.GenerateCalibrationAction(qr.Name))
        break;
      }
      case models.ActionConfig: {
        if qr.Configuration == nil {
          return c.String(http.StatusBadRequest, "configuration required")
        }
        actionCache.Push(actions.GenerateConfigAction(qr.Name, *qr.Configuration))
        break;
      }
      case models.ActionCode: {
        if qr.Code == nil {
          return c.String(http.StatusBadRequest, "code required")
        }
        actionCache.Push(actions.GenerateCodeResponse(qr.Name, *qr.Code))
        break;
      }
    }
    return c.String(http.StatusOK, "")
  }
}

func Setup(e *echo.Echo, conn *sql.DB, actionCache actions.Queue) {
  e.GET("/", home)
  e.GET("/status", ping)
  e.POST("/status", createStatus(conn, actionCache))
  e.POST("/queue", createQueue(actionCache))
}
