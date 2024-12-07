package routes

import (
	"database/sql"
	"fmt"
	"gardenometer/actions"
	"gardenometer/db"
	"gardenometer/email"
	"gardenometer/models"
	"gardenometer/status"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func createStatus(conn *sql.DB, actionCache *actions.Queue, emailClient *email.EmailClient) echo.HandlerFunc {
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
      log.Println(status.Err)
      emailClient.SendMail(fmt.Sprintf("Device %s delivered error: %s", status.Id, *status.Err))
      return c.String(http.StatusOK, "")
    }
    err = addRegistration(conn, &status)
    if err != nil {
      log.Println(err)
      return c.String(http.StatusInternalServerError, "registration error")
    }
    err = status.ValidateForInsert()
    if err != nil {
      log.Println(err)
      return c.String(http.StatusBadRequest, "must include all metric values")
    }
    metric := models.Metric{}
    metric.FromStatus(status)
    if err := db.InsertMetric(conn, metric); err != nil {
      log.Println(err)
      c.String(http.StatusInternalServerError, err.Error())
    }
    if err := applyAlerts(conn, metric, emailClient); err != nil {
      log.Println(err)
      return c.String(http.StatusInternalServerError, "alerts error")
    }
    a, found := actionCache.NextActionForName(status.Id)
    result := ""
    if found {
      result = a.String()
    }
    return c.String(http.StatusOK, result)
  }
}

func addRegistration(conn *sql.DB, status *models.Status) error {
  reg, err := db.ReadRegistration(conn, status.Id)
  if err != nil {
    return err
  }
  if reg == nil {
    newReq := models.NewRegistration(status.Id)
    err = db.InsertRegistration(conn, newReq)
    if err != nil {
      return err
    }
  }
  return nil
}

func applyAlerts(conn *sql.DB, metric models.Metric, emailClient *email.EmailClient) error {
  alerts, err := db.ReadAlertsForName(conn, metric.Name)
  if err != nil {
    return err
  }
  for _, alert := range alerts {
    switch alert.Key {
    case models.AlertMoisture:
      if metric.Moisture <= int(alert.Value){
        sendAlertEmail(metric.Id.String(),
          fmt.Sprintf("%d", metric.Moisture),
          fmt.Sprintf("%d", int(alert.Value)),
          emailClient)
      }
    case models.AlertTemp:
      if metric.Temp <= alert.Value {
        sendAlertEmail(metric.Id.String(),
          fmt.Sprintf("%f", metric.Temp),
          fmt.Sprintf("%f", alert.Value),
          emailClient)
      }
    }
  }
  return nil
}

func sendAlertEmail(id, actual, threshold string, emailClient *email.EmailClient) {
  emailClient.SendMail(
    fmt.Sprintf(
      "Device %s: Alert moisture level is %s which is at or below alert threshold %s",
      id, actual, threshold))
}
